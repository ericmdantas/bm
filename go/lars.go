package main

import (
	"fmt"
    "net/http"

    "github.com/go-playground/lars"
)

var (
	msg = []byte("yo!")
)

const (
	port = ":3000"
)

func main() {
    l := lars.New()
	
	fmt.Println(port)
    
    l.Get("/", func (c lars.Context) {
		c.Response().Write(msg)    
	})

    http.ListenAndServe(port, l.Serve())
}