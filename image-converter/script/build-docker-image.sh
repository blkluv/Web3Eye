#!/usr/bin/env bash
SHELL_FOLDER=$(cd "$(dirname "$0")";pwd)
PROJECT_FOLDER=$(cd $SHELL_FOLDER/../;pwd)
ROOT_FOLDER=$(cd $PROJECT_FOLDER/../;pwd)

set -o errexit
set -o nounset
set -o pipefail

configFile=$ROOT_FOLDER/config/config.toml
myConfigPath=$PROJECT_FOLDER/pkg/utils/
echo "sync config.toml from $configFile to $myConfigPath"
cp $configFile $myConfigPath
cp $configFile $myConfigPath

if git_status=$(git status --porcelain --untracked=no 2>/dev/null) && [[ -z "${git_status}" ]]; then
  git_tree_state=clean
fi

set +e
version=$(git describe --tags --abbrev=0)

if [ ! $? -eq 0 ]; then
  version=latest
fi
set -e

service_name=$(
  cd $PROJECT_FOLDER
  basename $(pwd)
)

## For development environment, pass the second variable
if [[ "x" != "x$1" ]]; then
  version=$1
fi

## For development environment, pass the second variable
if [[ "xdev" == "x$1" ]]; then
    version=latest
fi

# TODO: should be official registry
# registry=uhub.service.ucloud.cn
registry=""
OrginazeName=coastlinesss
# OrginazeName=web3eye

if [[ "x" != $2 ]]; then
  registry=$2/
fi

cd $PROJECT_FOLDER

user=$(whoami)
if [ "$user" == "root" ]; then
  docker build -t ${registry}${OrginazeName}/$service_name:$version .
else
  sudo docker build -t ${registry}${OrginazeName}/$service_name:$version .
fi
