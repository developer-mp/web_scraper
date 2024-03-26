package main

import (
	"log"

	"server/pkg/scrape"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	router := gin.Default()
	corsMiddleware := cors.Default()
	router.Use(func(c *gin.Context) {
		corsMiddleware.HandlerFunc(c.Writer, c.Request)
		c.Next()
	})
	router.POST("api/scrape", scrape.DisplayScrapingResults)
	router.POST("api/results", scrape.SaveScrapingResults)
	router.GET("api/results", scrape.GetScrapingResults)
	log.Fatal(router.Run(":8080"))
}