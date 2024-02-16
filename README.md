# Go GRPC Dev Container

This repository contains a Go GRPC example project that can be run in a dev container.

## Prerequisites

- [Install Docker](https://docs.docker.com/get-docker/)
- [Install VSCode](https://code.visualstudio.com/) with [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

## Getting Started

1. Clone this repository:
   ```shell
   $ git clone https://github.com/ericslandry/grpc.git
   ```
2. Open the repository in VSCode:
   ```shell
   $ code grpc
   ```
3. Reopen the repository in a container:
      - Click on the green icon in the bottom left corner of the window and select `Reopen in Container`.

4. Run the project:
   ```shell
   task
   ```
   You shouls see the following output:
   ```
   task: [build:proto] protoc *.proto --proto_path=. --go_out=. --go_opt=module=github.com/ericslandry/grpc/pb/greeter --go-grpc_out=. --go-grpc_opt=module=github.com/ericslandry/grpc/pb/greeter
   task: [setup] mkdir -p ./bin
   task: [build:client] go build -o ./bin/grpc-client ./cmd/client
   task: [setup] mkdir -p ./bin
   task: [build:server] go build -o ./bin/grpc-server ./cmd/server
   task: [run:server] ./bin/grpc-server &
   task: [run:server] until grpcurl -proto ./pb/greeter.proto -plaintext localhost:8080 list; do printf '.'; sleep 1; done
   grpc.greeter.v1.Greeter
   task: [run:client] ./bin/grpc-client --name=Mike
   2024/02/15 20:22:59 server listening at [::]:8080
   2024/02/15 20:22:59 Received: Mike
   2024/02/15 20:22:59 Greeting: Hello, Mike
   task: completed with code 0
   ```

## Testing
You can use [grpcurl](https://github.com/fullstorydev/grpcurl) to test the server:

```
$ grpcurl -proto ./pb/greeter.proto -plaintext localhost:8080 grpc.greeter.v1.Greeter/SayHello
{
  "greeting": "Hello, "
}
```
