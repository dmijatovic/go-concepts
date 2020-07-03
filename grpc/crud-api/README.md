# gRPC CRUD API example with MongoDB

This example uses Go and gRPC communication with MongoDB to demonstrate basic CRUD API. The example will be blog api enabling CRUD operations on blogposts.

## Project structure

- client: go client
- mongo: docker mongodb setup
- pb: grpc proto file and generated go file
- sever: go grpc server

## MongoDB

The mongoDB is installed using docker-compose container. To start mongoDB go to mongo folder and run docker-compose up. To close mongoDB run docker-compose down.

This demo uses [mongodb go driver](https://github.com/mongodb/mongo-go-driver). There is also go driver for mongo db.

## Generate protobuf go file

```bash
# get help on commands
protoc --help
# generate greet from greet folder
protoc ./pb/blog.proto --go_out=plugins=grpc:.
# this generates greet.pb.go file

```

## Dependecies

```bash
# install mongodb driver
go get github.com/mongodb/mongo-go-driver/mongo
# driver shows following message
# "go.mongodb.org/mongo-driver/mongo"

```
