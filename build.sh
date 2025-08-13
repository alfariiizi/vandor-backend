#!/bin/bash

set -e

APP_IMAGE_NAME="jogiia/jogiia:clinic-backend-1.0.0"
MIGRATE_IMAGE_NAME="jogiia/jogiia:clinic-migration-1.0.0"

build_app() {
  echo "🔧 Building app image..."
  docker buildx build --builder desktop-linux -t $APP_IMAGE_NAME -f docker/Dockerfile --platform=linux/amd64 .
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
