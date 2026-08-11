#!/usr/bin/env bash
set -Eeuo pipefail

# repair.sh for magic-dns
tailscale down
tailscale up --accept-dns=true
