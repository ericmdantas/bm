package main

import "github.com/kataras/iris/v12"

const (
	port = ":3000"
)

var (
	msg = []byte("yo!")
)

func main() {
	app := iris.New()
	app.Get("/", func(c iris.Context) {
		c.Write(msg)
	})

	app.Run(iris.Addr(port))
}
