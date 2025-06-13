# ---------- Stage 1: Build ----------
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server

#  ------------- Stage 2: Minimal Runtime ------------
FROM alpine:latest

RUN adduser -D appuser

WORKDIR /app

COPY --from=builder /app/server .

COPY --from=builder /app/Config ./Config

USER appuser

EXPOSE 8080

ENTRYPOINT ["./server"]
