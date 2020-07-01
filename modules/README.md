# Go modules

This folders contains experiments with Go mod command for creating packages/modules.

Most of the commands are learned from [this youtube video](https://youtube.com/watch?v=Z1VhG7cf83M)

When working with modules. The local imports are not possible. You do need to import using `<module-name>/<package-name>`, for example

## File imports

```go
// IN MODULE
// import local package token in the module dv4all-module-demo
import "dv4all-module-demo/token"

// WITHOUT MODULE SETUP
import "./token"

```

## Initialize go module project

```bash
# initialize
# this will create go.mod and state.go file
go mod init dv4all-module-demo

```

- main.go: contains module name, go version and 3rd party imports
- go.sum: contains dependecies graph for 3rd part libs

```bash
# get some thirt party dependencies (latest)
go get github.com/dgrijalva/jwt-go

# go has specific ways to handle 3rd party versions
# v1 and 2 are special. To load specific version use
# something like this
go get github.com/dgrijalva/jwt-go/v4

```

## Dockerfile

Building module app with Docker. This seem to work properly with 3rd parties

```bash
# build container
docker build . -t=dv4all/demo:v0.1

```
