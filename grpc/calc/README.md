# gRPC server/client calculator

This basic example is assignment from Udemy training. The goal is to create gRPC communication between client and server. The API to implement is SUM. The client sends two numbers which server adds and returns the result.

For basic implementatin see greet folder whic is basic hello world example for gRPC communication.

If this is your first project to try see root README file for installing basic gRPC dependecies on your PC.

## Dependencies

gRPC libraries just as in greet example. To use gRPC in go application go-grpc module is required.

```bash
# install go-grpc
go -u get google.golang.org/grpc
# protocol buffers support
go -u get github.com/golang/protobuf/protoc-gen-go

```

## Project stucture

- client: holds client gRPC go code
- pb: holds protocol buffers generate file based on `greet.proto` file definitions.
- server: holds gRPC server go code

## Compilation

For installation of protocol buffers see main README file. Here just code to create pb file.

```bash
# get help on commands
protoc --help
# generate calc protobuffer go file
protoc ./pb/calc.proto --go_out=plugins=grpc:.
```

## Running

```bash

# start sever from project root folder
go run ./server/main.go
# start client from project root folder
go run ./client/main.go

```
