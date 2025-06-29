# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Optional: install git for private modules
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app/main.go

# Stage 2: Runtime
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates bash curl

COPY --from=builder /app/main .
COPY --from=builder /app/storage ./storage

# Copy the entrypoint
COPY docker/entrypoint.sh .
RUN chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
