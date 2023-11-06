FROM caddy:builder-alpine AS builder
RUN xcaddy build \
  --with github.com/caddy-dns/cloudflare \
  --with github.com/caddyserver/replace-response \
  --with github.com/caddyserver/transform-encoder \
  --with github.com/corazawaf/coraza-caddy/v2

FROM caddy:alpine as runtime
COPY --from=builder /usr/bin/caddy /usr/bin/caddy
