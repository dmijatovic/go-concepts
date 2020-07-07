package handlers

import (
	"log"
	"net/http"

	"../data"
)

// Products creates handler
// which needs to implement interface
type Products struct {
	l *log.Logger
}

// NewProducts return pointer to a new instance of Products handler
// with desired logger.
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) logRequest(r *http.Request) {
	p.l.Println(r.URL.Path + "..." + r.Method)
}

func methodNotSupported(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	p.logRequest(req)
	switch req.Method {
	case http.MethodGet:
		getProducts(rw)
	case http.MethodPost:
		addProduct(rw, req)
	case http.MethodPut:
		updateProduct(rw, req)
	default:
		methodNotSupported(rw)
	}
}

func getProducts(rw http.ResponseWriter) {
	d := data.GetProducts()
	d.ToJSON(rw)
	return
}

func addProduct(rw http.ResponseWriter, req *http.Request) {
	// create new product struct
	prod := &data.Product{}
	// parse JSON body of the request
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(rw, "Unable to parse JSON", http.StatusBadRequest)
	}
	// log.Printf("%#v", prod)
	data.AddProduct(prod).ToJSON(rw)
	return
}

func updateProduct(rw http.ResponseWriter, req *http.Request) {
	// create new product struct
	prod := &data.Product{}
	// parse JSON body of the request
	err := prod.FromJSON(req.Body)
	// extract ID from url
	id := ExtractIDFromURL(req)
	prod.ID = id
	if err != nil {
		http.Error(rw, "Unable to parse JSON", http.StatusBadRequest)
	}
	// log.Printf("%#v", prod)
	data.UpdateProduct(prod).ToJSON(rw)
	return
}
