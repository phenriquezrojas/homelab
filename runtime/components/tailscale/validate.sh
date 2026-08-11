#!/usr/bin/env bash
set -Eeuo pipefail

# validate.sh for tailscale
if ! command -v tailscale >/dev/null 2>&1; then
    echo "ABSENT"
    exit 0
fi

if ! systemctl is-active --quiet tailscaled; then
    echo "INSTALLED"
    exit 0
fi

if ! tailscale status >/dev/null 2>&1; then
    echo "CONFIGURED"
    exit 0
fi

echo "HEALTHY"
exit 0
