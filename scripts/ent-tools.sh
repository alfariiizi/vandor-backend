#!/usr/bin/env bash

# ent-tools.sh — Lightweight utility for working with entgo

set -e

ENT_CMD="entgo.io/ent/cmd/ent"
SCHEMA_DIR="./database/schema"
GENERATED_DIR="internal/infrastructure/db"

function usage() {
  echo "Usage:"
  echo "  $0 new <SchemaName>         Create a new schema"
  echo "  $0 generate                 Generate Ent code"
  echo "  $0 install-deps             Install Ent CLI"
  echo "  $0 help                     Show this help"
  exit 1
}

function install_deps() {
  echo "Installing Ent CLI..."
  go install entgo.io/ent/cmd/ent@latest
  echo "Done. Make sure GOPATH/bin is in your PATH."
}

function create_schema() {
  if [ -z "$1" ]; then
    echo "Error: Missing schema name"
    exit 1
  fi

  go run -mod=mod "$ENT_CMD" new "$1" --target "$SCHEMA_DIR"
  echo "Schema '$1' created in $SCHEMA_DIR/$1.go"
}

function generate() {
  go run -mod=mod "$ENT_CMD" generate "$SCHEMA_DIR" --target "$GENERATED_DIR"
  echo "Code generation complete."
}

case "$1" in
  new)
    create_schema "$2"
    ;;
  generate)
    generate
    ;;
  install-deps)
    install_deps
    ;;
  help|*)
    usage
    ;;
esac
