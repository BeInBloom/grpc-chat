# 🔐 gRPC Chat

End-to-end encrypted chat application with gRPC-based microservices backend and native Rust client.

## Tech Stack

### Backend (Go)
- **gRPC** — service communication
- **Buf** — protobuf tooling
- **PostgreSQL** — data persistence
- **Docker** — containerization

### Client (Rust)
- Native desktop application
- End-to-end encryption

## Project Structure

```
.
├── proto/                  # Protobuf definitions
├── gen/go/                 # Generated Go code
├── services/
│   ├── auth/               # Authentication service
│   └── chat/               # Chat service
├── client/                 # Rust client
└── docker-compose.yml
```

## Quick Start

### Prerequisites
- Go 1.25+
- Docker & Docker Compose
- Buf CLI

### Run with Docker

```bash
docker compose up --build
```

### Local Development

```bash
# Install dependencies
make bin-deps

# Generate protobuf code
make generate

# Run auth service
go run ./services/auth/cmd

# Run chat service
go run ./services/chat/cmd
```

## Services

| Service | Port | Description |
|---------|------|-------------|
| auth    | 50051 | User authentication, JWT tokens |
| chat    | 50052 | Chat messaging, rooms |

## Development

### Git Hooks (Lefthook)

```bash
# Install hooks
lefthook install
```

Pre-commit: `goimports`, `golangci-lint`, `buf lint`  
Pre-push: `go test`, `go build`  
Commit-msg: Conventional Commits format

### Commit Message Format

```
type(scope): description

feat(auth): add JWT validation
fix(chat): resolve message ordering
chore: update dependencies
```

## License

MIT
