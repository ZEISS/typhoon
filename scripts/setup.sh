#!/bin/bash

set -x

# Get the current directory
dir=$(pwd)

# Create kind cluster
kind create cluster --config $dir/.cluster.yaml

# Knative serving version
knative_serving="1.17.0"

# Knative eventing version
knative_eventing="1.17.2"

# NATZ Accounting
natz_accounting="0.9.5"

# Setup helm charts
helm repo add natz-operator https://zeiss.github.io/natz-operator/helm/charts
helm repo add zeiss-staging https://zeiss.github.io/charts/staging
helm repo add nats https://nats-io.github.io/k8s/helm/charts
helm repo update

docker pull docker.io/envoyproxy/envoy:v1.31-latest
docker pull ghcr.io/zeiss/natz-operator/account-server:$natz_accounting
docker pull natsio/nats-box:0.16.0
docker pull nats:2.10.25-alpine
docker pull natsio/nats-server-config-reloader:0.16.1
docker pull ghcr.io/zeiss/natz-operator/operator:$natz_accounting
kind load docker-image docker.io/envoyproxy/envoy:v1.31-latest --name typhoon
kind load docker-image ghcr.io/zeiss/natz-operator/account-server:$natz_accounting --name typhoon
kind load docker-image natsio/nats-box:0.16.0 --name typhoon
kind load docker-image nats:2.10.25-alpine --name typhoon
kind load docker-image natsio/nats-server-config-reloader:0.16.1 --name typhoon
kind load docker-image ghcr.io/zeiss/natz-operator/operator:$natz_accounting --name typhoon

# # Install the Knative Serving and Eventing components on Minikube
helm install knative zeiss-staging/knative --wait

# # Install the Eventing components
helm install eventing zeiss-staging/eventing --wait

# See: https://gist.github.com/beriberikix/a827ec31f62705f13054895fa8cda0ad
kubectl apply --filename https://raw.githubusercontent.com/knative/serving/knative-v$knative_serving/third_party/kourier-latest/kourier.yaml

# Install the NATZ operator
helm install accounts natz-operator/natz-operator --wait
kubectl apply -f $(pwd)/examples/natz.yaml

# Install NATS.io
helm install sample-nats nats/nats --wait --values $(pwd)/examples/natz-server.yaml

# Install the account server
helm install account-server natz-operator/account-server --wait
kubectl apply -f $(pwd)/examples/natz-typhoon.yaml