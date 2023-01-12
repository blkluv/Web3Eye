#!/bin/bash
# install docs
# https://milvus.io/docs/install_cluster-helm.md

helm repo add milvus https://milvus-io.github.io/milvus-helm/
helm repo update

helm install milvus milvus/milvus
# have too many part,so wait more time
sleep 20
kubectl get pods | grep milvus

# helm install milvus milvus/milvus --set cpu=1
# helm install milvus milvus/milvus --set cluster.enabled=false --set etcd.replicaCount=1 --set minio.mode=standalone --set pulsar.enabled=false --set standalone.resources.limits.cpu: 3
