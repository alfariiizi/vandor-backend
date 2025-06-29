#!/bin/bash

# Shared configuration
MIGRATIONS_DIR="file://database/migrate/migrations"
ENT_SCHEMA_URL="ent://database/schema"
DEV_URL="docker://postgres/15/test?search_path=public"

# Choose how to load env
if [ -f .env ] && command -v dotenv >/dev/null 2>&1; then
  # Development: use dotenv-cli
  LOAD_ENV="dotenv -e .env --"
else
  # Production: assume env is already set
  LOAD_ENV=""
fi

COMMAND=$1
ARG=$2

case "$COMMAND" in
  diff)
    if [ -z "$ARG" ]; then
      echo "Usage: $0 diff <migration_name>"
      exit 1
    fi
    $LOAD_ENV atlas migrate diff "$ARG" \
      --dir "$MIGRATIONS_DIR" \
      --to "$ENT_SCHEMA_URL" \
      --dev-url "$DEV_URL"
    ;;
  apply)
    $LOAD_ENV sh -c "atlas migrate apply \
      --dir $MIGRATIONS_DIR \
      --url \$DB_URL"
    ;;
  status)
    $LOAD_ENV sh -c "atlas migrate status \
      --dir $MIGRATIONS_DIR \
      --url \$DB_URL"
    ;;
  *)
    echo "Usage: $0 {diff <name>|apply|status}"
    exit 1
    ;;
esac
