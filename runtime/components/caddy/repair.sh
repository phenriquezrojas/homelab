#!/usr/bin/env bash
set -Eeuo pipefail

# repair.sh for caddy
cd ../../services/caddy || { echo "services/caddy directory not found" >&2; exit 1; }
docker compose down
docker compose up -d
