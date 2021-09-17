package handlers

import (
	"log"
	"net/http"
	"github.com/4molybdenum2/microservices-arch/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts (l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle the request for a list of products
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// catch all


	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err!=nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}