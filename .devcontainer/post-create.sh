#!/bin/bash
set -ex

sudo apt-get update
sudo apt-get upgrade -y
# https://grpc.io/docs/protoc-installation/
sudo apt-get install -y protobuf-compiler
# https://protobuf.dev/getting-started/gotutorial/
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"
