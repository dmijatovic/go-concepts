package data

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// Product type defining json extraction too
type Product struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	CreatedOn string  `json:"-"`
	UpdatedOn string  `json:"-"`
	DeletedOn string  `json:"-"`
}

//Products is a collection
type Products []*Product

//GetProducts returns list of products
func GetProducts() Products {
	return productList
}

// AddProduct adds product to database
func AddProduct(p *Product) *Product {
	hashStr := getMd5HashStr(p)
	p.ID = hashStr
	p.CreatedOn = time.Now().UTC().String()
	productList = append(productList, p)
	return p
}

//UpdateProduct will update product in database
func UpdateProduct(p *Product) *Product {
	// here need to replace product

	// when done return product
	return p
}

//ToJSON returns json encoded products
func (p *Products) ToJSON(rw http.ResponseWriter) error {
	rw.Header().Set("content-type", "application-json")
	//return encoded products list
	return json.NewEncoder(rw).Encode(p)
}

//ToJSON returns json encoded product
func (p *Product) ToJSON(rw http.ResponseWriter) error {
	rw.Header().Set("content-type", "application-json")
	//return encoded products list
	return json.NewEncoder(rw).Encode(p)
}

//FromJSON will read data from body
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

var productList = Products{
	{
		ID:        "asdasde13",
		Name:      "Latte",
		Price:     2.45,
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	}, {
		ID:        "eradea2344",
		Name:      "Capuchino",
		Price:     3.51,
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
