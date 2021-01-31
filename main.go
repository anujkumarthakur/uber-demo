package main

import (
	//"github.com/gin-gonic/gin"
	"Uber/routes"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

//var router *gin.Engine
var client *redis.Client

func init() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the routes
	r := routes.InitializeRoutes()

	// Start serving the application
	//router.Run()
	r.Run(":" + os.Getenv("PORT"))

}
