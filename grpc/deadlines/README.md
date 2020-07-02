# Deadlines on gRPC request

Each request should define the dedadline. This example will perform simple request with defined deadline.

Context object is important for defining the deadlines in the request. The client who sends requests defines the deadline (how long he wants to wait). The server than checks the request context for the cancellation trigger/hook.

## Create PB go file

```bash
# get help on commands
protoc --help
# generate greet from greet folder
protoc ./pb/deadline.proto --go_out=plugins=grpc:.
# this generates greet.pb.go file

```

## Defining deadline on the client

We use context.WithTimeout fn as deadline hook. For complete implementation see client/main.go file.

```Go
// this defines the deadling of this request
ctx, cancel := context.WithTimeout(context.Background(), timeout)
// if cancel call cancel function
defer cancel()

```

## Checking cancellation on Server

The server should check from time to time if the request is cancelled. If so it should stop his internal operation related to this request.

Here is how we montor for cancel in this example. For complete implementation see server/main.go file method Task.

```Go

// we will wait for 3 seconds and check for
// cancelation each second
for i := 0; i < 3; i++ {
  // check if client cancelled request
  // using context (see client) for deadline implementation
  if ctx.Err() == context.Canceled {
    log.Println("Client cancelled request")
    return nil, status.Errorf(codes.Canceled, "The client cancelled request")
  }
  time.Sleep(1 * time.Second)
}

```
