package main

import (
	"fmt"
	"net/http"
)

const (
	port = ":3000"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("yo!"))
	})

	fmt.Println(port)

	http.ListenAndServe(port, nil)
}
