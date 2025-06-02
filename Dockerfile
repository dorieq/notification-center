FROM golang:1.24-alpine

WORKDIR /app

# Copy wait script
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build
RUN go build -o notify-service ./cmd/notify

# Start with dependency check
CMD ["/bin/sh", "-c", "/wait-for-it.sh kafka:9092 -- ./notify-service"]