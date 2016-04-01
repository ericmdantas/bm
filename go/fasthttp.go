package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

const (
	port = ":3000"
)

var (
	msg = []byte("yo!")
)

func main() {
	fmt.Println(port)

	fasthttp.ListenAndServe(port, func(ctx *fasthttp.RequestCtx) {
		ctx.Write(msg)
	})
}
