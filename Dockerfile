# syntax=docker/dockerfile:1

# Both runtime images are built for whatever --platform the caller asks for. The
# build stages pin themselves to the BUILD platform and cross-compile instead,
# so asking for linux/arm64 on an arm64 Mac, or linux/amd64 from the same Mac,
# both run at native speed rather than under emulation.

# ---- web -------------------------------------------------------------------
# The output is static files, identical on every architecture, so this always
# runs natively.
FROM --platform=$BUILDPLATFORM node:24-alpine AS web
WORKDIR /w
COPY apps/web/package.json apps/web/package-lock.json* ./
RUN npm ci || npm install
COPY apps/web/ ./
RUN npm run build

# ---- go --------------------------------------------------------------------
FROM --platform=$BUILDPLATFORM golang:1.27-alpine AS gobuild
ARG TARGETOS
ARG TARGETARCH
WORKDIR /src
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
COPY --from=web /w/dist ./apps/web/dist
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-arm64} \
    go build -trimpath -o /out/api ./apps/api \
 && CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-arm64} \
    go build -trimpath -o /out/stubgw ./apps/stubgw

# ---- runtime: our backend --------------------------------------------------
# ca-certificates is not in alpine's base rootfs. Nothing needs it while every
# URL is plain HTTP, and everything needs it the moment one is not: an RDS
# connection with sslmode=verify-full, or ALI_BASE_URL pointing at the real
# https://gw.open.1688.com. Adding it costs half a megabyte.
FROM alpine:3.21 AS api
RUN apk add --no-cache ca-certificates \
 && adduser -D -u 10001 app
COPY --from=gobuild /out/api /usr/local/bin/api
USER app
WORKDIR /home/app
EXPOSE 8787
ENTRYPOINT ["/usr/local/bin/api"]

# ---- runtime: the stub gateway ---------------------------------------------
# ALL-APIS.json is a hard startup dependency: the stub reads it to validate
# system parameters against the documented contract, and exits 1 if it is
# missing. Compose used to supply it as a bind mount, which meant the image only
# worked on a machine holding this repo. It is 275 KB; it belongs in the image.
FROM alpine:3.21 AS stub
RUN apk add --no-cache ca-certificates \
 && adduser -D -u 10001 app \
 && mkdir -p /docs /state \
 && chown app /state
COPY 1688-api-docs/ALL-APIS.json /docs/ALL-APIS.json
COPY --from=gobuild /out/stubgw /usr/local/bin/stubgw
USER app
WORKDIR /home/app
ENV DOCS_DIR=/docs \
    STUB_STATE_FILE=/state/stub-state.json
EXPOSE 8788
ENTRYPOINT ["/usr/local/bin/stubgw"]
