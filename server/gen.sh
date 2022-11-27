#!/bin/bash
set -euo pipefail

repo_root=$(git rev-parse --show-toplevel)
server_root="${repo_root}/server"
cd "${repo_root}"

rm -rf "${server_root}/pb"
mkdir "${server_root}/pb"
./protos/build.sh "${server_root}/pb" "github.com/blorente/gofood/server/pb"

