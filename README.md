# Go GRPC Dev Container

This repository contains a Go GRPC example project that can be run in a dev container.

## Prerequisites

- Docker: [Install Docker](https://docs.docker.com/get-docker/)
- VSCode: [Install VSCode](https://code.visualstudio.com/) with [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

## Getting Started

1. Clone this repository:
   ```shell
   $ git clone https://github.com/your-username/your-repo.git
   ```
2. Open the repository in VSCode:
   ```shell
   code your-repo
   ```
3. Reopen the repository in a container:
      - Click on the green icon in the bottom left corner of the window and select `Reopen in Container`.

4. Run the project:
   ```shell
   task
   ```
   You shouls see the following output:
   ```shell
   task: [build:proto] protoc *.proto --proto_path=. --go_out=. --go_opt=module=github.com/ericslandry/grpc/pb/greeter --go-grpc_out=. --go-grpc_opt=module=github.com/ericslandry/grpc/pb/greeter
   task: [setup] mkdir -p ./bin
   task: [build:client] go build -o ./bin/grpc-client ./cmd/client
   task: [setup] mkdir -p ./bin
   task: [build:server] go build -o ./bin/grpc-server ./cmd/server
   task: [run:server] ./bin/grpc-server &
   task: [run:client] ./bin/grpc-client --name=Mike
   2024/02/15 19:17:37 server listening at [::]:8080
   2024/02/15 19:17:37 Received: Mike
   2024/02/15 19:17:37 Greeting: Hello, Mike
   ```
