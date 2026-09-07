#!/usr/bin/env bash
# Deploy the marketplace to AWS: ECR, one Fargate task, an ALB and CloudFront.
#
#   ./deploy/deploy.sh            build, push and deploy everything
#   ./deploy/deploy.sh images     rebuild and push the images only
#   ./deploy/deploy.sh stack      redeploy the stack only
#   ./deploy/deploy.sh seed       import the catalogue
#   ./deploy/deploy.sh status     what is running, and where
#   ./deploy/deploy.sh logs       tail the task logs
#   ./deploy/deploy.sh destroy    delete the stack and the images
#
# The deploy runs in two passes on purpose. The stub stamps its own public
# origin into every product image URL, and those URLs are written into the
# database when the catalogue is imported. CloudFront's domain does not exist
# until the stack does, so the stack is created once to learn the domain and
# updated once to hand it back before anything is imported.

set -euo pipefail

STACK="${STACK:-marketplace}"
REGION="${AWS_REGION:-ap-southeast-1}"
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SECRET_NAME="${STACK}/config"
TAG="${TAG:-$(date -u +%Y%m%d-%H%M%S)}"

say()  { printf '\n\033[1m== %s\033[0m\n' "$*"; }
note() { printf '   %s\n' "$*"; }
die()  { printf '\033[31merror: %s\033[0m\n' "$*" >&2; exit 1; }

aws_() { aws --region "$REGION" "$@"; }

account_id() { aws_ sts get-caller-identity --query Account --output text; }
registry()   { echo "$(account_id).dkr.ecr.${REGION}.amazonaws.com"; }

stack_output() {
  aws_ cloudformation describe-stacks --stack-name "$STACK" \
    --query "Stacks[0].Outputs[?OutputKey=='$1'].OutputValue" --output text 2>/dev/null
}

# ---------------------------------------------------------------- credentials --

ensure_secret() {
  if aws_ secretsmanager describe-secret --secret-id "$SECRET_NAME" >/dev/null 2>&1; then
    note "secret $SECRET_NAME already exists, keeping the values it holds"
    return
  fi
  say "Generating credentials"
  local admin app_secret access control pgpass
  admin=$(openssl rand -hex 24)
  app_secret=$(openssl rand -hex 24)
  access=$(openssl rand -hex 24)
  control=$(openssl rand -hex 24)
  pgpass=$(openssl rand -hex 20)

  # DATABASE_URL carries the password, so it is generated here with the rest
  # rather than assembled in the task definition where a secret cannot be
  # interpolated into a string.
  aws_ secretsmanager create-secret --name "$SECRET_NAME" \
    --description "Marketplace runtime credentials" \
    --secret-string "$(cat <<JSON
{
  "ADMIN_TOKEN": "$admin",
  "APP_SECRET": "$app_secret",
  "ACCESS_TOKEN": "$access",
  "CONTROL_TOKEN": "$control",
  "POSTGRES_PASSWORD": "$pgpass",
  "DATABASE_URL": "postgres://app:$pgpass@localhost:5432/market?sslmode=disable"
}
JSON
)" >/dev/null
  note "created $SECRET_NAME with five generated credentials"
}

secret_arn()  { aws_ secretsmanager describe-secret --secret-id "$SECRET_NAME" --query ARN --output text; }
secret_key()  { aws_ secretsmanager get-secret-value --secret-id "$SECRET_NAME" --query SecretString --output text | python3 -c "import json,sys;print(json.load(sys.stdin)['$1'])"; }

# ---------------------------------------------------------------------- images --

build_and_push() {
  say "Building and pushing images (linux/arm64)"
  local reg; reg=$(registry)

  for repo in marketplace/api marketplace/stub; do
    aws_ ecr describe-repositories --repository-names "$repo" >/dev/null 2>&1 || {
      note "creating ECR repository $repo"
      aws_ ecr create-repository --repository-name "$repo" \
        --image-scanning-configuration scanOnPush=true >/dev/null
    }
  done

  aws_ ecr get-login-password | docker login --username AWS --password-stdin "$reg" >/dev/null
  note "logged in to $reg"

  # Fargate runs ARM64 here: it is native to an Apple Silicon build host, so
  # nothing is emulated, and it costs about 20% less than x86.
  docker buildx build --platform linux/arm64 --target api \
    -t "$reg/marketplace/api:$TAG" -t "$reg/marketplace/api:latest" --push "$ROOT"
  docker buildx build --platform linux/arm64 --target stub \
    -t "$reg/marketplace/stub:$TAG" -t "$reg/marketplace/stub:latest" --push "$ROOT"

  note "pushed :$TAG"
}

# ---------------------------------------------------------------------- stack --

deploy_stack() {
  local origin="${1:-}"
  local reg; reg=$(registry)
  local vpc subnets ip

  vpc=$(aws_ ec2 describe-vpcs --filters Name=isDefault,Values=true \
        --query 'Vpcs[0].VpcId' --output text)
  [ "$vpc" != "None" ] || die "no default VPC in $REGION"
  subnets=$(aws_ ec2 describe-subnets --filters "Name=vpc-id,Values=$vpc" \
            --query 'Subnets[0:3].SubnetId' --output text | tr '\t' ',')
  [ "$(echo "$subnets" | tr ',' '\n' | wc -l)" -ge 3 ] || die "need three subnets, found: $subnets"
  ip=$(curl -fsS https://checkip.amazonaws.com | tr -d '[:space:]')

  say "Deploying stack $STACK${origin:+ with origin $origin}"
  note "vpc $vpc"
  note "subnets $subnets"
  note "operator address $ip"

  aws_ cloudformation deploy \
    --stack-name "$STACK" \
    --template-file "$ROOT/deploy/cloudformation.yml" \
    --capabilities CAPABILITY_IAM \
    --no-fail-on-empty-changeset \
    --parameter-overrides \
      VpcId="$vpc" \
      SubnetIds="$subnets" \
      ApiImage="$reg/marketplace/api:$TAG" \
      StubImage="$reg/marketplace/stub:$TAG" \
      SecretArn="$(secret_arn)" \
      AdminIp="$ip" \
      AdminCidr="$ip/32" \
      PublicOrigin="$origin"
}

wait_healthy() {
  local url="$1" tries="${2:-60}"
  say "Waiting for $url/api/health"
  for i in $(seq 1 "$tries"); do
    if curl -fsS -m 5 "$url/api/health" 2>/dev/null | grep -q '"ok":true'; then
      note "healthy after ${i}0s or less"
      return 0
    fi
    sleep 10
  done
  note "still not healthy; recent logs:"
  logs 40
  return 1
}

# ----------------------------------------------------------------------- seed --

seed() {
  local url token
  url=$(stack_output PublicUrl)
  token=$(secret_key ADMIN_TOKEN)
  say "Importing the catalogue into $url"
  curl -fsS -m 30 -XPOST "$url/api/admin/import" \
    -H "Authorization: Bearer $token" -H 'Content-Type: application/json' \
    -d '{"all":true,"pages":8}' | head -c 200
  echo
  note "import queued; it takes a couple of minutes for ~300 products"
  for i in $(seq 1 40); do
    sleep 10
    local n
    n=$(curl -fsS -m 10 "$url/api/products?size=1" 2>/dev/null \
        | python3 -c 'import json,sys;print(json.load(sys.stdin).get("total",0))' 2>/dev/null || echo 0)
    printf '   %s products\r' "$n"
    [ "${n:-0}" -ge 280 ] && { echo; note "catalogue ready: $n products"; return 0; }
  done
  echo
  note "catalogue is still filling; check $url/api/products"
}

# --------------------------------------------------------------------- status --

status() {
  local url alb cluster service
  url=$(stack_output PublicUrl); alb=$(stack_output AlbUrl)
  cluster=$(stack_output ClusterName); service=$(stack_output ServiceName)
  say "Stack $STACK in $REGION"
  note "storefront  $url"
  note "admin       $url/admin"
  note "origin      $alb"
  [ -n "$cluster" ] && aws_ ecs describe-services --cluster "$cluster" --services "$service" \
    --query 'services[0].{desired:desiredCount,running:runningCount,pending:pendingCount,status:status,deployment:deployments[0].rolloutState}' \
    --output table
  echo
  note "admin token:   $(secret_key ADMIN_TOKEN)"
  note "control token: $(secret_key CONTROL_TOKEN)"
}

logs() {
  local group lines="${1:-60}"
  group=$(stack_output LogGroupName)
  [ -n "$group" ] || die "no log group yet"
  aws_ logs tail "$group" --since 15m --format short 2>/dev/null | tail -n "$lines"
}

destroy() {
  say "Deleting stack $STACK"
  aws_ cloudformation delete-stack --stack-name "$STACK"
  aws_ cloudformation wait stack-delete-complete --stack-name "$STACK" || true
  for repo in marketplace/api marketplace/stub; do
    aws_ ecr delete-repository --repository-name "$repo" --force >/dev/null 2>&1 || true
  done
  note "stack and images gone. The secret $SECRET_NAME is kept; delete it with:"
  note "aws secretsmanager delete-secret --secret-id $SECRET_NAME --force-delete-without-recovery --region $REGION"
}

# ----------------------------------------------------------------------- main --

case "${1:-all}" in
  images)  build_and_push ;;
  stack)   deploy_stack "$(stack_output PublicUrl)" ;;
  seed)    seed ;;
  status)  status ;;
  logs)    logs "${2:-60}" ;;
  destroy) destroy ;;
  all)
    command -v docker >/dev/null || die "docker is not installed"
    docker info >/dev/null 2>&1 || die "the docker daemon is not running"
    ensure_secret
    build_and_push
    deploy_stack ""                       # pass 1: learn the CloudFront domain
    url=$(stack_output PublicUrl)
    deploy_stack "$url"                   # pass 2: hand it back to the stub
    aws_ ecs update-service --cluster "$(stack_output ClusterName)" \
      --service "$(stack_output ServiceName)" --force-new-deployment >/dev/null
    aws_ ecs wait services-stable --cluster "$(stack_output ClusterName)" \
      --services "$(stack_output ServiceName)" || true
    wait_healthy "$url" || die "the service never became healthy"
    seed
    status
    ;;
  *) die "unknown command: $1" ;;
esac
