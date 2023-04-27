#!/usr/bin/env bash

# cd to the root path
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)"
cd "${ROOT}"

set -o errexit
set -o nounset
set -o pipefail

es_projects=(webui)

for proj in ${es_projects[@]}
do

cd ${ROOT}/$proj
if [[ ! -f .eslintrc.js ]]; then
    echo 'ERROR: missing .eslintrc.js in project root' >&2
    exit 1
fi

if ! command -v n; then
    npm install -g n -y
fi

if ! command -v yarn; then
    npm install -g yarn -y
fi


n v16.14.0
npm install @typescript-eslint/eslint-plugin@latest --save-dev
yarn lint

done



