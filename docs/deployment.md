# Deploying to AWS

One CloudFormation stack runs the whole product: the backend, the stub gateway and Postgres
in a single Fargate task, behind a load balancer and CloudFront. There is nothing to click.

```bash
./deploy/deploy.sh            # build, push, deploy and seed — about fifteen minutes
./deploy/deploy.sh status     # URLs and the generated tokens
./deploy/deploy.sh logs       # tail the running task
./deploy/deploy.sh destroy    # remove everything
```

`status` prints the storefront URL, the admin URL and the credentials that were generated on
first deploy. Nothing else is needed to use the deployment.

---

## What gets built

```
Browser ──HTTPS──► CloudFront  (free *.cloudfront.net certificate)
                     │  viewer-request function: 403 on /admin, /api/admin,
                     │  /cashier and /_control unless the viewer is the operator
                     ▼ HTTP
                   Application Load Balancer  (ingress limited to CloudFront)
                     ├─ /img/*, /cashier*  ──► stub  :8788
                     └─ everything else    ──► api   :8787
                     ▼
                   ECS Fargate · ARM64 · exactly one task
                     ├─ postgres   public.ecr.aws/docker/library/postgres:17-alpine
                     ├─ stub       ECR marketplace/stub
                     └─ api        ECR marketplace/api
                   EFS   Postgres data + stub order state
                   Secrets Manager   one secret, five generated credentials
```

### Why all three containers share one task

In `awsvpc` mode the containers of a task share a network namespace, so they reach each other
on `localhost` exactly as they do under Docker Compose. The backend finds Postgres on
`localhost:5432` and the gateway on `localhost:8788`; the stub pushes back to
`localhost:8787`. No service discovery, no internal load balancer, no cross-AZ traffic.

It also enforces something the code requires. The stub holds every order in memory behind a
mutex with no cross-process locking, so two instances would be incoherent: disjoint
catalogues, per-instance idempotency, and message ids that collide because each process seeds
them from its own clock. One task makes a second instance impossible. The service is set to
`MinimumHealthyPercent: 0` and `MaximumPercent: 100` for the same reason — a rolling deploy
stops the old task before starting the new one rather than briefly running both, which also
matters because two Postgres processes cannot share one data directory.

### Why the stub is on the public path

The stub serves the product images the storefront renders, and the cashier page that pays the
1688 side of an order. Both are browser-facing, so `/img/*` and `/cashier*` are routed to it.
Its `/_control/*` plane is never routed from the internet at all, and is separately guarded by
`STUB_CONTROL_TOKEN`. Reaching it goes through the backend's authenticated proxy at
`/api/admin/stub/*`, which needs both the admin token and the control token.

### Networking

The task runs in a public subnet with a public IP so it can pull from ECR through the internet
gateway. A private subnet would need a NAT gateway, which costs more than everything else in
this stack combined, and the application needs no outbound internet at all today: it talks
only to Postgres and the stub, both on localhost.

Three security groups, each admitting exactly one thing. CloudFront reaches the load balancer
(via the managed prefix list `pl-31a34658`, plus the operator's own address so the origin can
be checked before CloudFront propagates); the load balancer reaches the task; the task reaches
the filesystem.

---

## Prerequisites

* Docker running, with buildx. Images are built for `linux/arm64` and Fargate runs ARM64 —
  native on an Apple Silicon machine, and about 20% cheaper than x86.
* AWS credentials with permission to create ECR, ECS, EFS, ELB, CloudFront, IAM and Secrets
  Manager resources.
* A default VPC with three subnets in the target region. `AWS_REGION` selects the region and
  defaults to `ap-southeast-1`.

---

## The two-pass deploy

`./deploy/deploy.sh all` creates the stack twice on purpose.

The stub stamps its own public origin into every product image URL and into the cashier link,
and those URLs are **written into the database** when the catalogue is imported. CloudFront's
domain does not exist until the stack does. So the stack is created once to learn the domain,
updated once to hand it back to the stub, and only then is the catalogue imported.

Get this wrong and every image URL in the database points at `localhost`. The fix is not a
redeploy, it is a re-import.

Full sequence:

| # | Step | What happens |
|---|---|---|
| 1 | Secrets | Five random credentials into Secrets Manager as `<stack>/config`. Skipped if the secret already exists, so existing deployments keep working. |
| 2 | Images | ECR repositories created if missing, both images built for `linux/arm64` and pushed with a timestamp tag and `latest`. |
| 3 | Stack pass 1 | Everything, with `PublicOrigin` empty. |
| 4 | Stack pass 2 | Same stack with `PublicOrigin` set to the CloudFront domain, then a forced new deployment. |
| 5 | Wait | Polls `/api/health` until the backend reports its database is reachable. |
| 6 | Seed | Imports the stub catalogue, roughly 300 products, and waits for them to land. |

---

## Configuration

All five credentials are generated on first deploy and injected as ECS task **secrets**, never
as plain environment values, so they do not appear in the console or in
`describe-task-definition`.

| Key in the secret | Used by |
|---|---|
| `ADMIN_TOKEN` | the backend, for every `/api/admin/*` route |
| `APP_SECRET` | both sides: signs gateway requests and verifies inbound push messages |
| `ACCESS_TOKEN` | both sides: the only token the gateway accepts |
| `CONTROL_TOKEN` | the stub, guarding `/_control/*` |
| `POSTGRES_PASSWORD` | Postgres, and embedded in `DATABASE_URL` |
| `DATABASE_URL` | the backend. Generated with the password already in it, because a task definition cannot interpolate a secret into a string. |

The backend runs with `ENV=production`, which makes it **refuse to start** if any of those is
still a development placeholder such as `dev`, `devsecret` or `devtoken`. A deploy that forgets
one fails loudly instead of coming up wide open.

To rotate: edit the secret and force a new deployment.

```bash
aws secretsmanager update-secret --secret-id marketplace/config --secret-string '{...}'
aws ecs update-service --cluster marketplace --service marketplace --force-new-deployment
```

Note that `POSTGRES_PASSWORD` only takes effect on a fresh volume — Postgres sets the password
when it initialises its data directory, so changing it later also needs an `ALTER USER`.

---

## Access control

The storefront is public. The operator surface is not.

A CloudFront viewer-request function returns 403 for `/admin`, `/api/admin/*`, `/cashier` and
`/_control` unless the viewer's IP matches the address the deploy was run from. The load
balancer cannot do this: behind CloudFront every request arrives from a CloudFront address,
and the real client IP survives only in a header that anyone could set.

Your address changes, and then the admin console locks you out. Redeploy to pick up the new
one:

```bash
./deploy/deploy.sh stack
```

You can test the gate without redeploying, which is worth doing after any change to it:

```bash
FN=marketplace-admin-gate
ETAG=$(aws cloudfront describe-function --name $FN --query ETag --output text)
printf '{"version":"1.0","context":{"eventType":"viewer-request"},"viewer":{"ip":"1.2.3.4"},
"request":{"method":"GET","uri":"/api/admin/settings","headers":{},"querystring":{},"cookies":{}}}' > ev.json
aws cloudfront test-function --name $FN --if-match "$ETAG" --stage LIVE \
  --event-object fileb://ev.json --query 'TestResult.FunctionOutput' --output text
```

Beyond the gate, `/api/admin/*` still requires the bearer token, compared in constant time.

---

## Cost

Roughly **$40 a month** in `ap-southeast-1`, dominated by two fixed line items.

| Item | Monthly |
|---|---|
| Fargate ARM64, 0.5 vCPU + 2 GB, always on | ~$17 |
| Application Load Balancer | ~$17 |
| Public IPv4 addresses | ~$4 |
| EFS, a couple of GB | ~$1 |
| CloudFront, ECR, Secrets Manager at demo traffic | ~$1 |

Between demos, scale the task away and keep everything else standing:

```bash
aws ecs update-service --cluster marketplace --service marketplace --desired-count 0   # ~$18/mo
aws ecs update-service --cluster marketplace --service marketplace --desired-count 1
```

The load balancer is the item that buys least. It exists because CloudFront needs a stable
origin hostname and because `/img/*` has to reach a different container than `/`. A single EC2
instance running the Compose file behind Caddy would land near $15 for the same result, with
a machine to patch.

---

## Operations

```bash
./deploy/deploy.sh status                      # URLs, task state, tokens
./deploy/deploy.sh logs                        # last 15 minutes, all three containers
./deploy/deploy.sh images && ./deploy/deploy.sh stack   # ship a code change
aws logs tail /ecs/marketplace --follow --region ap-southeast-1
```

To reach the database, run a shell in the task:

```bash
TASK=$(aws ecs list-tasks --cluster marketplace --query 'taskArns[0]' --output text)
aws ecs execute-command --cluster marketplace --task "$TASK" --container postgres \
  --interactive --command "psql -U app -d market"
```

That needs `enableExecuteCommand` on the service, which is off by default; add it to the
service in `deploy/cloudformation.yml` if you want it.

---

## Verifying a deployment

```bash
U=$(aws cloudformation describe-stacks --stack-name marketplace \
    --query "Stacks[0].Outputs[?OutputKey=='PublicUrl'].OutputValue" --output text)
T=$(aws secretsmanager get-secret-value --secret-id marketplace/config \
    --query SecretString --output text | python3 -c 'import json,sys;print(json.load(sys.stdin)["ADMIN_TOKEN"])')

curl -s $U/api/health                                   # {"ok":true} — the backend reached Postgres
curl -s "$U/api/products?size=2" | grep -o 'cloudfront[^"]*' | head -2   # images on the CDN, not localhost
curl -s -o /dev/null -w '%{http_code}\n' $U/api/admin/settings -H "Authorization: Bearer $T"   # 200
curl -s -o /dev/null -w '%{http_code}\n' $U/api/admin/settings -H "Authorization: Bearer dev"  # 401

BASE_URL=$U ADMIN_TOKEN=$T python3 scripts/orderflow.py  # a full order, exit 0
```

The last one is the real test: it drives cart, the checkout preview gate, relay, the cashier
and tracking through to the consolidation warehouse against the live site, and exits non-zero
if the order stalls.

Two results that look wrong but are not. `/_control/state` returns **200 with HTML** — that is
the storefront's client-side routing catch-all, not the stub; the stub's control plane is
simply not routed. And `/api/health` returning 503 means Postgres is unreachable, not that the
process is down; the load balancer is deliberately slow to act on it, five failed checks, so a
slow query cannot cycle the task.

---

## Troubleshooting

**The task will not start.** Read the stopped-task reason first, it is usually explicit:

```bash
aws ecs describe-tasks --cluster marketplace \
  --tasks $(aws ecs list-tasks --cluster marketplace --desired-status STOPPED --query 'taskArns[0]' --output text) \
  --query 'tasks[0].{reason:stoppedReason,containers:containers[].{name:name,reason:reason,exit:exitCode}}'
```

**The backend exits immediately** with a message about development values. `ENV=production` is
doing its job: one of the credentials in the secret is still a placeholder.

**The stub exits with `read /docs/ALL-APIS.json`.** The image was built without it. That file
is a hard startup dependency and is copied in by the `stub` stage of the Dockerfile; check
`.dockerignore` still has the `!1688-api-docs/ALL-APIS.json` negation.

**`exec format error`.** The images were built for the wrong architecture. Fargate here is
ARM64; build with `--platform linux/arm64`.

**Postgres will not initialise.** The EFS access point sets uid and gid 70, which is what
`postgres:17-alpine` runs as (the Debian images use 999), and mode 0700, which Postgres
insists on for its data directory. If the access point is changed, match those.

**Health check retries.** ECS caps container health check retries at 10. First boot has to run
`initdb` against EFS, which is slower than local disk, so give it time through `StartPeriod`
rather than the retry count.

---

## Teardown

```bash
./deploy/deploy.sh destroy
```

This deletes the stack, and with it the EFS filesystem and every order in the database, then
deletes both ECR repositories. The Secrets Manager entry is kept deliberately, because
deleting a secret is slow to reverse; remove it explicitly when you are sure:

```bash
aws secretsmanager delete-secret --secret-id marketplace/config --force-delete-without-recovery
```

---

## When real 1688 access arrives

The deployment changes are small, because the backend already speaks the real protocol.

* Point `ALI_BASE_URL` at `https://gw.open.1688.com` and put the real app key, secret and
  token in the secret.
* Give the task outbound internet: a NAT gateway, or a VPC endpoint. Today it needs none.
* The runtime images already carry CA certificates, so TLS will work.
* Register the push URL with 1688 and adapt `internal/api/hooks.go` to their real signature
  contract, replacing the invented one.
* Delete the stub container, its two load balancer rules, the `/api/admin/stub/*` proxy and
  the admin Stub tab.

The code-side checklist is in the main README under *Going live*; verifying
`internal/ali/sign.go` against the official documentation is the first item on it, because the
stub validates with the same function and cannot tell you it is wrong.
