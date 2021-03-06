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

func setupRouter() *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))
	r.GET("/", handler.RequestToTomTom)
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
