package main

import (
	"fmt"
	"net/http"
)

const (
	port = ":3000"	
)

var (
	msg = []byte("yo!")
)


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {	
		w.Write(msg)
	})

	fmt.Println(port)

	http.ListenAndServe(port, nil)
}
