#!/usr/bin/env bash
set -Eeuo pipefail

# repair.sh for tailscale
systemctl restart tailscaled
