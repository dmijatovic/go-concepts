# Error in gRPC

There are handful err codes used by gRPC protocol. The list is [avaliable here](https://grpc.io/docs/guides/error.html)

To add more info to error you can use error context.

There is extended [guide for handling errors in gRPC here](https://grpc.io/docs/guides/error.html).

There is also example implementation at [this website](https://avi.im/grpc-errors/)

In this example we implement INVALID_VALUE error code.

## Create PB go file

```bash
# get help on commands
protoc --help
# generate greet from greet folder
protoc ./pb/errors.proto --go_out=plugins=grpc:.
# this generates greet.pb.go file

```

## Dependencies

To be able to return types gRPC error we need to import 2 additional grpc modules.

```Go

import (
  "google.golang.org/grpc/status"
  "google.golang.org/grpc/codes"
)

```
