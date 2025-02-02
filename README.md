# typhoon

[![Test & Lint](https://github.com/ZEISS/typhoon/actions/workflows/main.yml/badge.svg)](https://github.com/ZEISS/typhoon/actions/workflows/main.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/zeiss/typhoon)](https://goreportcard.com/report/github.com/zeiss/typhoon)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)
[![Volkswagen](https://auchenberg.github.io/volkswagen/volkswargen_ci.svg?v=1)](https://github.com/auchenberg/volkswagen)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

An event bridge for applications and services build on an event mesh with a team-oriented self-service control plane.

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/ZEISS/typhoon?quickstart=1)

## Design

Every software system has design goals. Properties that it should provide io its users to create value. Typhoon has the following design goals:

* **Team-based self-service control plane**. Teams control their own event sources, triggers, and targets. They also control their users their permissions.
* **Event mesh**. The system is capable to route events from producers to consumers in hybrid cloud environemts.
* **Streaming**. The system is capable to stream events from producers to consumers. That means it is capable of replaying events not only queueing them.
* **Event bridging**. The system is capable to bridge events from different sources to different targets, also transforming them. Furthermore, this should use a declarative approach.
* **Scalability**. The system is capable to scale horizontally and vertically.
* **Resilience**. The system is capable to recover from failures.
* **Security**. The system is capable to secure the communication between producers and consumers.
* **Open standards**. The system is based on open standards like [NATS.io](https://nats.io/) and [Knative Eventing](https://knative.dev/docs/eventing/).

## Overview

Typhoon is built on top of [NATS.io](https://nats.io/) and [Knative Eventing](https://knative.dev/docs/eventing/). It provides a control plane for managing event sources, triggers, and targets. It also provides an API for managing the control plane.

```mermaid
flowchart TB

subgraph controlPlane [Control Plane]
  subgraph api[API]
    subgraph apiAccounting[Accounting]
    end
  end
end

subgraph "Knative Eventing"
  subgraph triggers[Triggers]
    subgraph sources[Sources]
    end
    subgraph transformers[Transformers]
    end
    subgraph targets[Targets]
    end
  end
end

triggers--Communicates via-->eventMesh[Event Mesh]
apiAccounting--Stores data-->externalDatabase[External Database]
apiAccounting--Provides authentication-->eventMesh[Event Mesh]
controlPlane-->ssoProvider[SSO Provider]

sources-->transformers
transformers-->targets

subgraph eventMesh[Event Mesh]
  subgraph NATS
  end
end

subgraph ssoProvider[SSO Provider]
end

subgraph externalDatabase[External Database]
  subgraph PostgreSQL
  end
end
```

The accounting can also be combined with the [Operator for NATS Accounting](https://github.com/ZEISS/natz-operator) which enables the creation of operators, accounts and users declaratively.

## Sources

* [CloudEvents](https://cloudevents.io/)
* HTTP Poller
* [Kafka](https://kafka.apache.org/)
* [Salesforce](https://www.salesforce.com/)
* WebHook

## Targets

* [CloudEvents](https://cloudevents.io/)
* DataDog
* JIRA
* [NATS](https://nats.io)
* [Salesforce](https://www.salesforce.com/)
* [ServiceNOW](https://www.servicenow.com/)
* Splunk

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

## Prerequisites

* Kubernetes `v1.28` or newer
* Knative Eventing `v1.15` or newer
* Knative Serving `v1.15` or newer
* Helm `3.0` or newer

The prerequisites can be installed via the [ZEISS Charts](https://github.com/ZEISS/charts).

## Installation

Typhoon depends on [Knative Eventing](https://knative.dev/docs/) and [NATS](https://nats.io).

## Development

You can use [minikube](https://minikube.sigs.k8s.io/docs/) to run a local Kubernetes cluster.

```shell
sh scripts/setup-minikube.sh
```

Run the following Makefile target to deploy Typhoon to your local Kubernetes cluster.

```shell
make deploy
```

## License

[Apache 2.0](/LICENSE)
