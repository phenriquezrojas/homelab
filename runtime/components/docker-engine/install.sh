#!/usr/bin/env bash
set -Eeuo pipefail

# install.sh for docker-engine
if ! command -v curl >/dev/null 2>&1; then
    apt-get update && apt-get install -y curl
fi

curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh
rm get-docker.sh
