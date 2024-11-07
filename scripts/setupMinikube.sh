#!/bin/bash

set -x

# Get the current directory
dir=$(pwd)

# Install Minikube
minikube start --cpus=2 --memory=4096 --addons=ingress

# Knative serving version
knative_serving="1.16.0"

# Knative eventing version
knative_eventing="1.16.0"

# Setup helm charts
helm repo add natz-operator https://zeiss.github.io/natz-operator/helm/charts
helm repo add zeiss-staging https://zeiss.github.io/charts/staging
helm repo add nats https://nats-io.github.io/k8s/helm/charts
helm repo update

# Install the Knative Serving and Eventing components on Minikube
helm install knative zeiss-staging/knative --wait --create-namespace

# Install the Eventing components
helm install eventing zeiss-staging/eventing --wait

# See: https://gist.github.com/beriberikix/a827ec31f62705f13054895fa8cda0ad
kubectl apply --filename https://raw.githubusercontent.com/knative/serving/knative-v$knative_serving/third_party/kourier-latest/kourier.yaml

# Install the NATZ operator
helm install natz-operator natz-operator/natz-operator --wait --namespace knative-eventing

# Create operator resources
kubectl apply -f $(pwd)/examples/natz-operator.yaml

# Create account resources
kubectl apply -f $(pwd)/examples/natz-account.yaml

# Create user resources
kubectl apply -f $(pwd)/examples/natz-user.yaml

# Install NATS.io
helm install nats nats/nats --wait --values $(pwd)/examples/nats-server.yaml

# Install the NATZ accounts-server
helm install account-server natz-operator/account-server --wait --namespace knative-eventing --values $(pwd)/examples/natz-account-server.yaml

# MiniKube IP
minikube tunnel -c
