#!/bin/bash

set -e

APP_IMAGE_NAME="alfariiizi/go-app:latest"
MIGRATE_IMAGE_NAME="alfariiizi/go-migrate:latest"

build_app() {
  echo "🔧 Building app image..."
  docker buildx build -t $APP_IMAGE_NAME -f docker/Dockerfile --platform=linux/amd64 .
}

build_migrate() {
  echo "🔧 Building migration image..."
  docker buildx build -t $MIGRATE_IMAGE_NAME -f docker/Dockerfile.migrate --platform=linux/amd64 .
}

case "$1" in
  app)
    build_app
    ;;
  migrate)
    build_migrate
    ;;
  all)
    build_app
    build_migrate
    ;;
  *)
    echo "Usage: $0 {app|migrate|all}"
    exit 1
    ;;
esac
