#!/bin/bash

set -x

# Install Knative Serving
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.13.1/serving-crds.yaml
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.13.1/serving-core.yaml

# Install Networking Layer
kubectl apply -f https://github.com/knative/net-kourier/releases/download/knative-v1.13.0/kourier.yaml
kubectl patch configmap/config-network \
  --namespace knative-serving \
  --type merge \
  --patch '{"data":{"ingress-class":"kourier.ingress.networking.knative.dev"}}'
kubectl --namespace kourier-system get service kourier
kubectl get pods -n knative-serving

# Configure Magic DNS
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.13.1/serving-default-domain.yaml

# Install Knative Eventing
kubectl apply -f https://github.com/knative/eventing/releases/download/knative-v1.13.3/eventing-crds.yaml
kubectl apply -f https://github.com/knative/eventing/releases/download/knative-v1.13.3/eventing-core.yaml

# Install NATS JetStream
kubectl apply -f example/knative-eventing-jetstream-crds.yaml
kubectl apply -f https://github.com/knative-extensions/eventing-natss/releases/latest/download/eventing-jsm.yaml
kubectl apply -f example/knative-nats.yaml
kubectl apply -f example/knative-eventing-config-nats.yaml
kubectl apply -f example/knative-eventing-default-channel.yaml