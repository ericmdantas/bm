package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	port = ":3000"
)

func main() {
	fmt.Println(port)

	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "", "yo!")
	})

	router.Run(port)
}
