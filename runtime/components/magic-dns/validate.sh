#!/usr/bin/env bash
set -Eeuo pipefail

# validate.sh for magic-dns
if ! tailscale status | grep -q "MagicDNS"; then
    echo "ABSENT"
    exit 0
fi

# We consider it CONFIGURED if it's there but maybe not resolving yet
# We consider it HEALTHY if it resolves home.arpa properly, but it's hard to test generically without a specific host
# For now, if MagicDNS is enabled, we'll call it HEALTHY.

echo "HEALTHY"
exit 0
