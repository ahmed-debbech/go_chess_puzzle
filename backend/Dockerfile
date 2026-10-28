
# Build stage
FROM golang:1.23.2-alpine AS builder

RUN apk add --no-cache git ca-certificates

COPY . .

WORKDIR ./backend

RUN go get .


RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /out/backend .

# Runtime stage
FROM alpine:3.22

RUN apk add --no-cache ca-certificates \
    && adduser -D -H -u 10001 appuser

WORKDIR /app

COPY --from=builder /out/backend ./backend

USER appuser

EXPOSE 5530

ENTRYPOINT ["./backend"]
