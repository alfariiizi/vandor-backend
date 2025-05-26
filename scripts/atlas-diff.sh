#!/usr/bin/env bash

set -e

# === Config ===
MODELS_SQL="database/models.sql"
MIGRATION_DIR="database/migrations"
DEV_URL="postgres://root:root@localhost:5432/sqlc-testing?sslmode=disable"

# === Load .env file ===
if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
fi

# === Validate DB_URL is set ===
if [ -z "$DEV_URL" ]; then
  echo "❌ Error: DB_URL not found. Please set it in your .env file."
  exit 1
fi

# === Require a migration name argument ===
if [ -z "$1" ]; then
  echo "❌ Usage: $0 \"migration name\""
  exit 1
fi

# === Run Atlas migration diff ===
atlas migrate diff "$1" \
  --to "file://$MODELS_SQL" \
  --dev-url "$DEV_URL" \

echo "✅ Migration generated in: $MIGRATION_DIR"
