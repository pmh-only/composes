FROM alpine AS builder

RUN apk add go

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

COPY main.go /app/

RUN go build main.go

FROM scratch AS runtime

WORKDIR /app

COPY --from=builder /lib/ld-musl-x86_64.so.1 /lib/ld-musl-x86_64.so.1
COPY --from=builder /app/main /app/

CMD ["/app/main"]
