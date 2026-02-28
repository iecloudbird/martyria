# Build stage
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /martyria ./cmd/martyria

# Runtime stage
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata

# Create non-root user
RUN addgroup -S martyria && adduser -S martyria -G martyria

WORKDIR /app

COPY --from=builder /martyria /app/martyria
COPY migrations/ /app/migrations/

# Create data directories
RUN mkdir -p /app/data/images && chown -R martyria:martyria /app

USER martyria

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD wget -qO- http://localhost:8080/health || exit 1

ENTRYPOINT ["/app/martyria"]
