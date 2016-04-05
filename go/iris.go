package main

import (
	"fmt"

	"github.com/kataras/iris"
)

const (
	port = ":3000"
)

func main() {
	fmt.Println(port)

    iris.Get("/", func(c *iris.Context) {
        c.Text("yo!")
    })
	
    iris.Listen(port)
}