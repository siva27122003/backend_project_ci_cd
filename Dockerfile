# ---------- Stage 1: Build ----------
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# Copy Go mod files 
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

RUN go build -o server

#  ------------- Stage 2: Minimal Runtime ------------
FROM alpine:latest

# Create non-root user
RUN adduser -D appuser

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/server .

# Set permissions (optional but safe)
USER appuser

# App port (change if your app uses another)
EXPOSE 8080

# Run the app
ENTRYPOINT ["./server"]
