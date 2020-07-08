# http Servers in Go

This section is based on [excellent youtube training from Nic Jackson about building Microservices](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW).

The training starts with very basic http setup and builds on to a proffesional grade micrososervice.

In this readme I create notes and learning which are specific to my observations. I will create few different sub projects as progressing.

## Leranings

### Video 1: Basics of http

- Use any reader to read body of request. For displaying erros use http.Error object

```Go
// read
body,err := ioutil.ReadAll(r.Body)
if err != nil {
  http.Error(rw, "Bad request", http.StatusBadRequest)
}

```

See basic1 folder for implementation.

### Video 2: Custom handlers and decoupling for more flexibility

The goal is to extract handler and enable flexible logging and be able to test.
We create http server and customize the settings.

```Go
//create server
api := &http.Server{
  Addr: host,
  // Handler can be anything that
  // implements ServeHTTP method
  Handler:      router,
  IdleTimeout:  60 * time.Second,
  ReadTimeout:  10 * time.Second,
  WriteTimeout: 10 * time.Second,
}
```

Gracefuly closing server is also important and can be implemented like this.

```Go
// register for os signals
// listening for closing server
stop := make(chan os.Signal)
// broadcast message
signal.Notify(stop, os.Interrupt)
signal.Notify(stop, os.Kill)

// wait here to receive os signals
sig := <-stop

// close all systems
l.Println("Terminate signal received:", sig)
l.Println("Starting gracefull shutdown...")
//define waiting for grafull shutdown to 30 secs
tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
api.Shutdown(tc)
```

### Video 3&4: Separating data source from route handlers and JSON implementation

This secion shows basic of json implementation and separation of data from handlers. The data related code is placed in its own package `data`.There is also ToJSON method implemented that will take care of json response.

Convert struct data to JSON

```Go
//ToJSON returns json encoded products
func (p *Products) ToJSON(w io.Writer) error {
 //return encoded products list
 return json.NewEncoder(w).Encode(p)
}
```

Read JSON body and pass it into struct. Note that request Body has implementation of io.Reader and therefore we can pass it to FromJSON method.

```Go
// product HANDLER module
// create new product struct
prod := &data.Product{}
// parse JSON body of the request
err := prod.FromJSON(req.Body)
```

```Go
// product DATA module
//FromJSON will read data from body
func (p *Product) FromJSON(r io.Reader) error {
 e := json.NewDecoder(r)
 return e.Decode(p)
}

```

For exact implementation see basis3 folder.

### Video 5: using gorilla mux instead of standard one

In this video we learn how to use [gorilla/mux router](https://github.com/gorilla/mux).

First we need to get package.

```bash
# get gorilla mux router
go get github.com/gorilla/mux
```

Gorilla supports varable routes. See products/id handler.

In addition it also supports middleware functions. The middleware has next function which returns handler

### Video 6: using validator package

For validating posted data we can use [validator module](https://pkg.go.dev/github.com/go-playground/validator/v10?tab=doc).

See video and repo for the implementation.

### Video 7&8: swagger and redoc

In these video Nic show how to implement [swagger](https://goswagger.io/use/spec.html) and redoc docs. Swagger and redoc modules need to be installed.

Another online tutorial for [swagger-go can be found here](https://medium.com/@pedram.esmaeeli/generate-swagger-specification-from-go-source-code-648615f7b9d9)

### Video 10: uploading files

In the folder uploads there is a demo for accepting files upload using POST method and binary body. The image files are uploaded to storage folder.

The demo handles two types of uploads. POST and MULTIPART form.

For detailed implementation see uploads folder.
