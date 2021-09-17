package data

import (
	"io"
	"encoding/json"
	"time"
)

// structure of the product our API will expose
type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	SKU string `json:"sku"`
	Price float32 `json:"price"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

// Products is a collection of Product
type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts () Products {
	return productList
}

// productList is a hard coded list of products for this
// example data source
var productList = Products{
	&Product{
		ID:          1,
		Name:        "Latte",
		Desc: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Desc: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}