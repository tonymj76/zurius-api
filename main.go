package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tonymj76/zurius-api/handler"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", handler.GooglePlace)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
