#!/bin/bash
set -euo pipefail

outdir="$1"
go_module_root="$2"
repo_root=$(git rev-parse --show-toplevel)
cd "${repo_root}"
protos_root="${repo_root}/protos"
set -x
files=$(find ${protos_root} -type f -name "*.proto" | sed "s:${protos_root}/::")
# TODO: TIdy this to avoid duplicating these sed calls
go_module_mappings=$(echo "${files}" | \
  sed 's:\(.*\)\.proto:--go_opt=M\1.proto=__go_module_root__/\1:' | \
  sed "s:__go_module_root__:${go_module_root}:")
grpc_module_mappings=$(echo "${files}" | \
  sed 's:\(.*\)\.proto:--go-grpc_opt=M\1.proto=__go_module_root__/\1:' | \
  sed "s:__go_module_root__:${go_module_root}:")
protoc \
    --proto_path="${protos_root}" \
    --go_out="${outdir}" \
    --go_opt=paths=source_relative \
    "${go_module_mappings[@]}" \
    --go-grpc_out="${outdir}" \
    --go-grpc_opt=paths=source_relative \
    "${grpc_module_mappings[@]}" \
    "${files[@]}"

