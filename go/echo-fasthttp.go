package main

/*
build github.com/ericmdantas/bm/go: cannot load github.com/labstack/echo/v4/engine/fasthttp:
  module github.com/labstack/echo/v4@latest
  found (v4.1.11), but does not contain package github.com/labstack/echo/v4/engine/fasthttp


import (
	"net/http"
	"fmt"

	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo"
)

const (
	port = ":3000"
)

func main() {
	e := echo.New()	
	
	fmt.Println(port)
	
	e.Get("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "yo!")
	})
	
	e.Run(fasthttp.New(port))
}
*/