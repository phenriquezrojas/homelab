#!/usr/bin/env bash
set -Eeuo pipefail

# validate.sh for caddy
if ! docker image inspect caddy:alpine >/dev/null 2>&1; then
    echo "ABSENT"
    exit 0
fi

if ! docker ps -q -f name=^caddy$ >/dev/null 2>&1; then
    echo "INSTALLED"
    exit 0
fi

# Signature check: Is it actually responding to our health check?
# Our Caddyfile explicitly returns "OK" for health.home.arpa
RESPONSE=$(curl -s -H "Host: health.home.arpa" http://localhost:80 || true)
if [ "$RESPONSE" = "OK" ]; then
    echo "HEALTHY"
    exit 0
fi

# If it didn't respond "OK", is something else occupying the port?
if lsof -Pi :80 -sTCP:LISTEN -t >/dev/null 2>&1; then
    echo "FAILED"
    exit 0
fi

echo "CONFIGURED"
exit 0
