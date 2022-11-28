#!/bin/bash
set -euo pipefail

repo_root=$(git rev-parse --show-toplevel)
cd "${repo_root}"

output="$HOME/.cache/gofood/data"
mkdir -p "${output}"
cp -r "${repo_root}/server/pb_data" "${output}"
