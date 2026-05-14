# --- Build Stage ---
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

# Leverage Docker cache for dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build a statically linked binary
# -ldflags="-w -s" reduces binary size by removing debug info
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o /app/api-server cmd/main.go

# --- Final Stage ---
FROM alpine:3.21

# Security: Don't run as root
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Install essential runtime utilities
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /home/appuser

# Copy only the binary from the builder
COPY --from=builder /app/api-server .

# Copy environment example if needed, but prefer ENV variables
# COPY .env.example .env

# Change ownership to non-root user
RUN chown appuser:appgroup api-server

# Use non-root user
USER appuser

# Expose the application port
EXPOSE 8080

# Healthcheck to ensure the container is healthy
# Note: Requires curl or a custom healthcheck tool in alpine
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/users || exit 1

ENTRYPOINT ["./api-server"]
