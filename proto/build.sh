#!/bin/bash
set -euo pipefail

repo_root=$(git rev-parse --show-toplevel)
cd "${repo_root}"
protoc \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --proto_path=../ \
    --go-grpc_opt=paths=source_relative \
    proto/*.proto
