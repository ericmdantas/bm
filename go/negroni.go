package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
)

const (
	port = ":3000"
)

var (
	msg = []byte("yo!")
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(msg)
	})

	fmt.Println(port)

	n := negroni.New()
	n.UseHandler(mux)

	n.Run(port)
}
