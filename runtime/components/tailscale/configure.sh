#!/usr/bin/env bash
set -Eeuo pipefail

# configure.sh for tailscale
systemctl enable --now tailscaled

# If we are already connected, nothing to do
if tailscale status >/dev/null 2>&1; then
    exit 0
fi

echo "Action Required: Tailscale is not authenticated." >&2
echo "Please run 'tailscale up' manually to authenticate, then re-run homelab converge." >&2
exit 1
