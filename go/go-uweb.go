package main

import (
	"fmt"
	"net/http"
	
	uweb "github.com/ericmdantas/go-uweb"
)

const port = ":3000"

var msg = []byte("yo!")

func main() {
	u := uweb.New()
	
	fmt.Println(port)
	
	u.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(msg)
	})
	
	http.ListenAndServe(port, u)
}