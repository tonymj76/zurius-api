package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tonymj76/zurius-api/handler"
)

// init gets called before the main function
func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

var url = "https://zurius-client.herokuapp.com"

// var urlx = "http://localhost:3000"

func setupRouter() *gin.Engine {
	r := gin.Default()
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:  []string{urlx},
	// 	AllowMethods:  []string{"GET", "HEAD", "OPTIONS"},
	// 	AllowHeaders:  []string{"Origin"},
	// 	ExposeHeaders: []string{"Content-Type"},
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == urlx
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{urlx}
	r.Use(cors.New(config))
	r.GET("/", handler.RequestToTomTom)
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	return r
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT must be set")
	}
	r := setupRouter()
	r.Run(":" + port)
}
