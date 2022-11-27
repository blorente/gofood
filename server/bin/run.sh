#!/bin/bash
set -euo pipefail

repo_root=$(git rev-parse --show-toplevel)
server_root="${repo_root}/server"

cd "${server_root}"
go run github.com/blorente/gofood/server "${@}"

