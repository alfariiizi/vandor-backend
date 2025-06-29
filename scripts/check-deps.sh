#!/bin/bash

# Color helpers
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

# List of dependencies to check
REQUIRED_TOOLS=(
  "go"
  "atlas"
  "dotenv"
)

echo "🔍 Checking required local development dependencies..."

MISSING=0

for TOOL in "${REQUIRED_TOOLS[@]}"; do
  if ! command -v "$TOOL" >/dev/null 2>&1; then
    echo -e "${RED}✗ $TOOL is not installed${NC}"
    MISSING=1
  else
    VERSION=$($TOOL version 2>/dev/null | head -n1 || echo "✓ found")
    echo -e "${GREEN}✓ $TOOL is installed${NC} — $VERSION"
  fi
done

if [ $MISSING -eq 1 ]; then
  echo -e "\n${RED}Some dependencies are missing. Please install them before continuing.${NC}"
  exit 1
else
  echo -e "\n${GREEN}All dependencies are installed!${NC}"
fi
