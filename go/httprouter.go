package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	port = ":3000"
)

var (
	msg = []byte("yo!")
)

func main() {
	router := httprouter.New()

	fmt.Println(port)

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write(msg)
	})

	http.ListenAndServe(port, router)
}
