FROM golang:1.25-alpine AS builder

WORKDIR /build

COPY go.work go.work.sum ./
COPY gen/go/go.mod gen/go/go.sum ./gen/go/
COPY pkg/go.mod ./pkg/
COPY services/auth/go.mod services/auth/go.sum ./services/auth/
COPY services/chat/go.mod ./services/chat/

RUN go mod download

COPY gen/go ./gen/go
COPY pkg ./pkg
COPY services/auth ./services/auth

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/auth ./services/auth/cmd

FROM alpine:3.21

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/auth /app/auth

EXPOSE 50051

ENTRYPOINT ["/app/auth"]
