package main

import (
	"log"
	"os"
	"time"

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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://zurius-client.herokuapp.com", "http://localhost:3000/"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://zurius-client.herokuapp.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/api/v1", handler.RequestToTomTom)

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
