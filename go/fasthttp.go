package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

const (
	port = ":3000"
)

func main() {
	fmt.Println(port)

	fasthttp.ListenAndServe(port, func(ctx *fasthttp.RequestCtx) {
		ctx.Write([]byte("yo!"))
	})
}
