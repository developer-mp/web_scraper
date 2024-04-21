package main

import (
	"log"

	"server/pkg/redisdb"
	"server/pkg/scrape"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "",
    DB:       0,
    })
    defer redisClient.Close()

	redisdb.InitializeClient(redisClient)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:8081"},
    AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
	AllowHeaders:     []string{"Content-Type"},
  }))

	v1 := router.Group("/api/v1")
	{
		v1.POST("/scrape", scrape.DisplayScrapingResults)
		v1.POST("/results", scrape.SaveScrapingResults)
		v1.GET("/results", scrape.GetScrapingResults)
		v1.DELETE("/results/:id", scrape.DeleteScrapingResults)
	}
	
	log.Fatal(router.Run(":8080"))
}