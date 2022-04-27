package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Running Gin implementation on http://localhost:4000")

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "World")
	})
	r.Run(":4000")
}
