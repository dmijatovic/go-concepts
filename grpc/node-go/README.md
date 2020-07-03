# gRPC communication between NodeJS client and Go server

This demo used nodejs on client side and golang on server side to communicate over gRPC

## Generate PB scripts

```bash
# get help on commands
protoc --help
# generate calc protobuffer go file
protoc ./pb/calc.proto --go_out=plugins=grpc:.
# generate nodejs
protoc ./pb/calc.proto --go_out=plugins=grpc:.

```
