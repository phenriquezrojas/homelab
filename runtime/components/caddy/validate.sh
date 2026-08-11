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

if ! curl -fsSL -o /dev/null http://localhost:80; then
    echo "CONFIGURED"
    exit 0
fi

echo "HEALTHY"
exit 0
