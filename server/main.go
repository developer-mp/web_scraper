package main

import (
	"log"

	"server/pkg/scrape"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:8081"},
    AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
  }))

	router.POST("api/scrape", scrape.DisplayScrapingResults)
	router.POST("api/results", scrape.SaveScrapingResults)
	router.GET("api/results", scrape.GetScrapingResults)
	router.DELETE("api/results/:id", scrape.DeleteScrapingResult)
	log.Fatal(router.Run(":8080"))
}