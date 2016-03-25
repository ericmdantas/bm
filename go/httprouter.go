package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	port = ":3000"
)

func main() {
	router := httprouter.New()

	fmt.Println(port)

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("yo!"))
	})

	http.ListenAndServe(port, router)
}
