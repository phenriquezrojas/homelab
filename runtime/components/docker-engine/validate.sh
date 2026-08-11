#!/usr/bin/env bash
set -Eeuo pipefail

# validate.sh for docker-engine
if ! command -v docker >/dev/null 2>&1; then
    echo "ABSENT"
    exit 0
fi

if ! systemctl is-active --quiet docker; then
    echo "INSTALLED"
    exit 0
fi

if ! docker info >/dev/null 2>&1; then
    echo "CONFIGURED"
    exit 0
fi

echo "HEALTHY"
exit 0
