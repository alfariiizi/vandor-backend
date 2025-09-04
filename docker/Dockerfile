# ====== Builder ======
# Use Alpine-based Go image for a small builder.
ARG GOLANG_VERSION=1.24
FROM golang:${GOLANG_VERSION}-alpine AS builder

# build deps only in builder
RUN apk add --no-cache git ca-certificates

WORKDIR /src

# cache go modules layer
COPY go.mod go.sum ./
RUN go mod download

# copy rest
COPY . .

# build outputs to /out for easy copy
ARG APP_VERSION=dev
ARG BUILD_TIME=unknown
ARG GOOS=linux
ARG GOARCH=amd64

ENV CGO_ENABLED=0
# Use BuildKit cache mounts if available (fast incremental builds)
# Requires BuildKit (DOCKER_BUILDKIT=1). If not using BuildKit, remove the --mount lines.
RUN --mount=type=cache,target=/root/.cache/go-build \
	--mount=type=cache,target=/go/pkg/mod \
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 \
	go build -trimpath -ldflags="-s -w -X 'main.version=${APP_VERSION}' -X 'main.buildTime=${BUILD_TIME}'" \
	-o /out/main ./cmd/app/main.go

# ====== Final (scratch) ======
FROM scratch AS final

# copy CA bundle from the builder (so HTTPS works)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# copy the built binary
COPY --from=builder /out/main /app/main

# If you have static assets in /app/storage needed at build time, copy them; otherwise prefer a PVC.
# COPY --from=builder /src/storage /app/storage

# run as non-root uid (no /etc/passwd in scratch; k8s will still see UID)
USER 1000

WORKDIR /app

# expose your port (optional)
# EXPOSE 8080

# binary must be the entrypoint in scratch
ENTRYPOINT ["/app/main"]
