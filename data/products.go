package data

import (
	"fmt"
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

var ErrProductNotFound = fmt.Errorf("Product not found")

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProducts () Products {
	return productList
}

func AddProduct (p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}


func UpdateProduct(id int, p*Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}
func getNextID () int {
	lp := productList[len(productList) - 1]
	return lp.ID + 1
}


func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
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