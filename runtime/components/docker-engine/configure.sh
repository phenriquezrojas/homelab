#!/usr/bin/env bash
set -Eeuo pipefail

# configure.sh for docker-engine
systemctl enable --now docker
