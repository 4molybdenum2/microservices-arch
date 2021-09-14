package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is a simple handler
type Hello struct {
	l *log.Logger
}

// sayHello creates a new hello handler with the given logger
func SayHello(l *log.Logger) *Hello {
	return &Hello{l}
}
func(h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World !")
	d,err := ioutil.ReadAll((r.Body))
	if err!=nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s", d)
}