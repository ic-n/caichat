# Caichat

This repository provides a template for building a fully self-hosted service for creating a Chat AI-assisted Software as a Service (SaaS) application. The service is written in Go and leverages the Ollama framework (https://github.com/jmorganca/ollama) to enable easy setup and configuration of large language models locally. This template exposes a gRPC HTTP API for customization and provides Kubernetes configuration managed by Helm, allowing you to easily deploy and scale your service.

## Prerequisites

Before you begin, make sure you have the following prerequisites:

- Docker installed to build and package the service.
- Helm (https://helm.sh/) installed for Kubernetes deployment and management.
- kubectl (Kubernetes command-line tool) configured to interact with your Kubernetes cluster.
- Go for development installed (https://go.dev/doc/install)
- Buf for managing protobuf files installed (https://buf.build/docs/installation)

## Getting Started

To get started with building and deploying your Chat AI-assisted SaaS using this template, follow the steps below:

### Build

1. Build the Docker image for your service by running the following command:

   ```sh
   task build
   ```

   This command will build a Docker image tagged with a random identifier.

### Deploy

2. Deploy your service using Helm by running the following command:

   ```sh
   task deploy
   ```

   This will install or upgrade your service in your Kubernetes cluster, using the Docker image built in the previous step.

Your service is now up and running and accessible through the configured Kubernetes resources.

## Uninstall

If you need to uninstall your service, use the following command:

```sh
task uninstall
```

This command will uninstall the service from your Kubernetes cluster.

## Monitoring and Logs

You can monitor your service and access logs as follows:

### Logs

To access logs for your service, use the following command:

```sh
task logs
```

### Port Forward

If you want to access your application locally, you can use port forwarding with the following command:

```sh
task port-forward
```

This will forward the service's port to your local machine, allowing you to interact with your application at `http://127.0.0.1:8080`.

### Development Code Generation

Run the following commands to lint and generate code for your project:

```sh
task dev-gen
```

These commands will ensure that your protocol buffers are in good shape by linter and generate any required code.

## Development

- protofiles defining your API is located [/proto/caichat/v1](https://github.com/ic-n/caichat/tree/main/proto/caichat/v1)
- service implementation located at [/pkg/api/api.go](https://github.com/ic-n/caichat/blob/main/pkg/api/api.go)
- helm configuration located in [/deploy](https://github.com/ic-n/caichat/tree/main/deploy) as well as [dockerfile](https://github.com/ic-n/caichat/blob/main/deploy/caichat.dockerfile)
