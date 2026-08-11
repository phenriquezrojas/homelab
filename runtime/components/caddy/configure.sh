#!/usr/bin/env bash
set -Eeuo pipefail

# configure.sh for caddy

if lsof -i :80 | grep -q apache2; then
    echo "Port 80 is in use by apache2. Please disable it before proceeding." >&2
    exit 1
fi

mkdir -p /srv/homelab/app_data/caddy/data
mkdir -p /srv/homelab/app_data/caddy/config

# We assume Caddyfile is in services/caddy/Caddyfile as per previous structure.
# But since this is executed from the repo root or runtime, we should reference it relative to project root.
# Let's just run docker compose from services/caddy.
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
cd "$SCRIPT_DIR/../../../services/caddy" || { echo "services/caddy directory not found" >&2; exit 1; }
docker compose up -d

# Wait for it to be running
for i in {1..10}; do
    if docker ps -q -f name=^caddy$ >/dev/null; then
        exit 0
    fi
    sleep 1
done

echo "Caddy container failed to start" >&2
exit 1
