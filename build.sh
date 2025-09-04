#!/usr/bin/env bash

set -euo pipefail

APP_NAME="analisis_produk"
APP_VERSION="0.4.3"
REGISTRY="jogiia/jogiia"

APP_IMAGE_NAME="${REGISTRY}:${APP_NAME}-be-${APP_VERSION}"
MIGRATE_IMAGE_NAME="${REGISTRY}:${APP_NAME}-migrator-${APP_VERSION}"
TOOLS_IMAGE_NAME="${REGISTRY}:${APP_NAME}-tools-${APP_VERSION}"

BUILDER="${BUILDER:-desktop-linux}"
PLATFORMS="${PLATFORMS:-linux/amd64}"

build_app() {
  echo "🔧 Building app image with BuildKit (version: $APP_VERSION)..."
  docker buildx build \
    --builder "${BUILDER}" \
    --platform "${PLATFORMS}" \
    -t "${APP_IMAGE_NAME}" \
    -f docker/Dockerfile \
    --load .
}

build_migrate() {
  echo "🔧 Building migrate image with BuildKit (version: $APP_VERSION)..."
  docker buildx build \
    --builder "${BUILDER}" \
    --platform "${PLATFORMS}" \
    -t "${MIGRATE_IMAGE_NAME}" \
    -f docker/Dockerfile.migrate \
    --load .
}

show_help() {
  cat <<EOF
Usage: $0 <command> [push]

Commands:
  app         build app image (docker/Dockerfile)
  migrate     build migration image (docker/Dockerfile.migrate)
  tools       build tools image (docker/Dockerfile.tools)
  all         build app + migrate + tools
  clean       remove local images by tag (if present)

If you supply the second argument 'push' (or --push), the script will perform buildx --push,
which publishes the image to the registry and does NOT keep a local image.
Examples:
  $0 app
  $0 app push
  $0 all push
  $0 clean    # remove local images by tag (if present)
EOF
}

build_tools() {
  echo "🔧 Building tools image (fat) (version: $APP_VERSION)..."
  docker build -t "${TOOLS_IMAGE_NAME}" -f docker/Dockerfile.tools .
}

buildx_check() {
  if ! docker buildx ls >/dev/null 2>&1; then
    echo "⚠️  docker buildx not available. Please enable buildx or install Docker Buildx."
    exit 1
  fi
}

clean_local() {
  set +e
  echo "🧹 Scaning image with version: $APP_VERSION..."
  echo "🧹 Cleaning local images (if present)"
  docker image rm "${APP_IMAGE_NAME}" "${MIGRATE_IMAGE_NAME}" "${TOOLS_IMAGE_NAME}" >/dev/null 2>&1 || true
  set -e
}

case "${1:-}" in
  app)
    buildx_check
    build_app
    ;;
  migrate)
    buildx_check
    build_migrate
    ;;
  tools)
    build_tools
    ;;
  all)
    buildx_check
    build_app
    build_migrate
    build_tools
    ;;
  clean)
    clean_local
    ;;
  help|*)
    show_help
    ;;
esac
