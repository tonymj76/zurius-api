package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tonymj76/zurius-api/handler"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://zurius-client.herokuapp.com"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://zurius-client.herokuapp.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/", handler.GooglePlace)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
