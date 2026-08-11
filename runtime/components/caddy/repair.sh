#!/usr/bin/env bash
set -Eeuo pipefail

# repair.sh for caddy
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
cd "$SCRIPT_DIR/../../../services/caddy" || { echo "services/caddy directory not found" >&2; exit 1; }
docker compose down
docker compose up -d
