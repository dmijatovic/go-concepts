# Basic api demo with GO and Postgres

This demo project uses basis Go database/sql module and official pg driver for postgres.

The intention is to create simple api with less as possible dependencies, so we use:

- basic go database module and postgres driver
- basic go http router (defaul mux)
- encode/json for json marshaling
- bcrypt module for password hashing

## Module stucture

There are some ideas about MVC structure but I am inclined using a custom structure:

- hash: module resposible for hashing the passwords. bcrypt is used.
- pgdb: module responsible for postgres database connection and models
- routes module: container routes and calls appropriate method in models
- views: static index.html page for root of the api

## Dependencies

This modules need to be installed if you do not have them already

```bash
# pg driver
go get github.com/lib/pq
# bcrypt
go get golang.org/x/crypto/bcrypt

```

## Go quirks

- json backticks definitions need to be enclosed with "" and no spaces between `json:` and value.

```go
// ServerStatus is send as part of api response
// this is quite similair how axios responds
type ServerStatus struct {
 Status     int    `json:"status"`
 StatusText string `json:"statusText"`
}

```

- json ignore field use `json:"-"`
- private/Public methods, functions and properties in Go are defined by first letter. If it is Capital then method or property is Public, if it start with smallCase it is private.
- Response Headers can be added with method Add. This need to be done BEFORE setting return status code using WriteHeader method. If not the Header value will not be written back to consumer!!!

```go
// ReturnResponse will return response to api consumer
// including the status code
// NOTE! When setting header values, this need to be
// done before setting status using WriteHeader!!!
func (r *Response) ReturnResponse(rw http.ResponseWriter) {
 // set content-type
 rw.Header().Add("content-type", "application/json")
 rw.Header().Add("x-server", "dv4all-basic-go-http-server")
 // SET ALL HEADER PROPS BEFORE setting state
 // NOTE! GO requirement
 rw.WriteHeader(r.Status)
 // log.Println(rw)
 json.NewEncoder(rw).Encode(r)
}
```
