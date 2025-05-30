# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git (needed to download dependencies)
RUN apk add --no-cache git

# Copy go.mod and go.sum and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN go build -o notify-service ./cmd/notify

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/notify-service .

# Optional: add CA certs if needed
RUN apk add --no-cache ca-certificates
