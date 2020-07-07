package handlers

import (
	"log"
	"net/http"

	"../data"
	"github.com/gorilla/mux"
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

// func (p *Products) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
// 	p.logRequest(req)
// 	switch req.Method {
// 	case http.MethodGet:
// 		getProducts(rw)
// 	case http.MethodPost:
// 		addProduct(rw, req)
// 	case http.MethodPut:
// 		updateProduct(rw, req)
// 	default:
// 		methodNotSupported(rw)
// 	}
// }

//GetProducts returns array of products
func GetProducts(rw http.ResponseWriter, req *http.Request) {
	d := data.GetProducts()
	d.ToJSON(rw)
	return
}

// AddProduct to database
func AddProduct(rw http.ResponseWriter, req *http.Request) {
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

// UpdateProduct will update product by id
func UpdateProduct(rw http.ResponseWriter, req *http.Request) {
	//extract id
	vars := mux.Vars(req)
	// create new product struct
	prod := &data.Product{}
	// parse JSON body of the request
	err := prod.FromJSON(req.Body)
	// extract ID from url
	prod.ID = vars["id"]
	if err != nil {
		http.Error(rw, "Unable to parse JSON", http.StatusBadRequest)
	}
	// log.Printf("%#v", prod)
	data.UpdateProduct(prod).ToJSON(rw)
	return
}
