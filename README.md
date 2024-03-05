# typhoon

![Github Action](https://github.com/zeiss/typhoon/workflows/main/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/zeiss/typhoon.svg)](https://pkg.go.dev/github.com/zeiss/typhoon)
[![Go Report Card](https://goreportcard.com/badge/github.com/zeiss/typhoon)](https://goreportcard.com/report/github.com/zeiss/typhoon)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)
[![Volkswagen](https://auchenberg.github.io/volkswagen/volkswargen_ci.svg?v=1)](https://github.com/auchenberg/volkswagen)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

An event bridge for applications and services.

## Helm

[Helm](https://helm.sh/) can be used to install Typhoon to your Kubernetes cluster.

```shell
helm repo add typhoon https://zeiss.github.io/typhoon
helm repo update
```

Install Typhoon to your cluster in a `typhoon` namespace.

```shell
helm install typhoon typhoon/typhoon --create-namespace --namespace typhoon
```

## Development

You can use [minikube](https://minikube.sigs.k8s.io/docs/) to run a local Kubernetes cluster.

```shell
minikube start
```

You need to install Serving and Eventing for Knative.

## License

[Apache 2.0](/LICENSE)