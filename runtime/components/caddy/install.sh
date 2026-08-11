#!/usr/bin/env bash
set -Eeuo pipefail

# install.sh for caddy
docker pull caddy:alpine
docker network inspect homelab-net >/dev/null 2>&1 || docker network create homelab-net
