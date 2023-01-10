#!/bin/bash
SHELL_FOLDER=$(
    cd "$(dirname "$0")"
    pwd
)
ROOT_FOLDER=$(
    cd $SHELL_FOLDER/../
    pwd
)

registry=""
OrginazeName=coastlinesss
# OrginazeName=web3eye

user=$(whoami)
service_name=jenkins
version=latest
if [ "$user" == "root" ]; then
    docker build -t ${registry}${OrginazeName}/$service_name:$version .
else
    sudo docker build -t ${registry}${OrginazeName}/$service_name:$version .
fi
