FROM alpine

RUN apk add --no-cache nodejs
RUN addgroup -S sandbox && adduser -S sandbox -G sandbox

USER sandbox:sandbox

WORKDIR /home/sandbox
