#!/usr/bin/env bash
set -Eeuo pipefail

# repair.sh for docker-engine
systemctl restart docker
