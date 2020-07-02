# SSL encription with gRPC

This demo uses self signed certificates. The complete process is described below.

Official gRPC guide for TLS implemenation [is here](https://grpc.io/docs/guides/auth.html).

And example implementation is [available here](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-auth-support.md).

For SSL encription we will crete following files

- `ca.key` (private - server): Certificate Authority Private Key (not for sharing)
- `ca.crt` (share - client): Certificate Authority Trust Certificate
- `server.key` (private - server): Server private key, password protected
- `server.csr` (share - CA): Server Certificate Signing Request (to be shared with CA owner)
- `server.crt` (private-server): Server Certificate Signed by CA. This file is send back by CA owner and reside on server
- `server.pem` (private - server): Converted server.key to gRPC required format

The files shared with others are:

- ca.crt for the client
- server.csr send to CA

## SSL generation (self-signed approach)

These steps can be performed on Linux or other machines that have openssl installed. The example will use SERVER_CN bash variable for localhost.

These commands should be executed in cert folder. All steps can be set in bash script.

```bash
#!/bin/bash

SERVER_CN=localhost

# 1. Generate ca.key and ca.crt file (Certificate Authority and Trust Certificate)

openssl genrsa -passout pass:1111 -des3 -out ca.key 4096
openssl req -passin pass:1111 -new -x509 -days 365 -key ca.key -out ca.crt -subj "/CN=${SERVER_CN}"

# 2. Generate Server Private Key (server.key)
openssl genrsa -passout pass:1111 -des3 -out server.key 4096

# 3. Create signed request by CA (server.csr)
openssl req -passin pass:1111 -new -key server.key -out server.csr -subj "/CN=${SERVER_CN}"

# 4. Sign certificate with CA we created (self-signing), server.crt
openssl x509 -req -passin pass:1111 -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt

# 5. Convert server certificate to .pem format (server.pem)
openssl pkcs8 -topk8 -nocrypt -passin pass:1111 -in server.key -out server.pem

```

## Go code adjustments for SSL/TSL connections

On the servier side we need to point to certFile and keyFile and start new server with TLS. For exact implementation see server/main.go methods newGRPCTLSrv

```Go
var certFile string = "/home/dusan/test/go/grpc/ssl/cert/server.crt"
var keyFile string = "/home/dusan/test/go/grpc/ssl/cert/server.pem"

// SSL credentials setup
creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
if sslErr != nil {
  log.Fatalln("Failed to load SSL...", sslErr)
}
// add SSL credentials to options
opts := grpc.Creds(creds)
// create grpc server with SSL options
s := grpc.NewServer(opts)

```

On the client side following need to be adjusted
