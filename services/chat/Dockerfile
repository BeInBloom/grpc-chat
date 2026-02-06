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
COPY services/chat ./services/chat

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/chat ./services/chat/cmd

FROM alpine:3.21

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/chat /app/chat

EXPOSE 50052

ENTRYPOINT ["/app/chat"]
