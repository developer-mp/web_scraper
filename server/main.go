package main

import (
	"log"

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
	router.POST("/scrape", handleScrape)
	log.Fatal(router.Run(":8080"))
}