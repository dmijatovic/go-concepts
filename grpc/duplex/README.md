# gRPC bi-directional streams (full-duplex)

This example implements full-duplex stream between client and server.

Still we have server/client setup, meaning the server can close connection/stop client stream at any time.

## Build protobuffer file

```bash
# get help on commands
protoc --help
# generate greet from greet folder
protoc ./pb/duplex.proto --go_out=plugins=grpc:.
# this generates greet.pb.go file

```

## Chat service PB

The proto file defines a ChatService which is bi-directional stream. For all details see pb/duplex.proto file.

## Server side implementation

For this duplex example the server needs to implement Chat method according to pb interface defined in pb/duplex.pb.go. Chat methods receives ChatService_ChatServer interface with Send and Recv methods.

## Client side implementation

IMPORTANT: at the client side, both sending and receiving messages NEED to be from one created `stream` object in order to have bi-directional communication. I made mistake and created two stream objects which did not worked properly.

On the client side we could also implement go routines for parallel processing
