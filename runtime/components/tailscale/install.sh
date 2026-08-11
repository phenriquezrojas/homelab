#!/usr/bin/env bash
set -Eeuo pipefail

# install.sh for tailscale
if ! command -v curl >/dev/null 2>&1; then
    apt-get update && apt-get install -y curl
fi

curl -fsSL https://tailscale.com/install.sh | sh
