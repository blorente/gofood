#!/bin/bash
set -euo pipefail

outdir="$1"
repo_root=$(git rev-parse --show-toplevel)
cd "${repo_root}"
protos_root="${repo_root}/protos"
protoc \
    --proto_path="${protos_root}" \
    --go_out="${outdir}" \
    --go_opt=paths=source_relative \
    --go-grpc_out="${outdir}" \
    --go-grpc_opt=paths=source_relative \
    "${protos_root}/*.proto"
