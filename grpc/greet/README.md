# Unary gRPC communication setup

This is basic gRPC setup with server and client exchanging the messages in the `unary` way. This is basic communication setup using protocol buffers. The connection is not secured in this demo.

## Project stucture

- client: holds client gRPC go code
- pb: holds protocol buffers generate file based on `greet.proto` file definitions.
- server: holds gRPC server go code

## Create pb.go service

For installation of protocol buffers see main README file. Here just code to create pb file.

```bash
# get help on commands
protoc --help
# generate greet from greet folder
protoc ./pb/greet.proto --go_out=plugins=grpc:.
# this generates greet.pb.go file

# or run shell
bash create_pb.sh
```

## To start server and client

Each of the instances should be started in separate bash window. Server will stay awake after start.

```bash

# start sever from greet folder
go run ./server/main.go
# start client from greet folder
go run ./client/main.go

```

## Proto files

In the pb/greet.proto file we define service and methods avaliable.

## Server learnings

Server needs to implement GreetServiceServer interface defined in ./pb/greet.pb.go file.

```go
// server interface definition greet.pb.go
// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
 //Unary API approach
 // maing convention RootName (RootNameRequest) (RootNameResponse)
 Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

```
