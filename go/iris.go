package main

import (
	"fmt"

	"github.com/kataras/iris"
)

const (
	port = ":3000"
)

func main() {
	app := iris.New()
	app.Get("/", func(c iris.Context) {
		c.Text("yo!")
	})

	app.Run(iris.Addr(port))
}
