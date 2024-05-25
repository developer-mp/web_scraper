package main

import (
	"os"
	logger "server/internal"
	"server/pkg/redisdb"
	"server/pkg/scrape"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/sebest/logrusly"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

var logglyToken string
var host string
var clientPort string
var serverPort string
var redisPort string

// Rate limiter allowing 10 requests per second with bursts of up to 5 requests
var limiter = rate.NewLimiter(10, 5)

func rateLimitMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        if !limiter.Allow() {
            c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
            return
        }
        c.Next()
    }
}

func main() {
	l := logrus.New()
	h := logrusly.NewLogglyHook(logglyToken, host, logrus.WarnLevel, "backend")
	l.Hooks.Add(h)

    err := godotenv.Load(".env.development")
    if err != nil {
        l.Fatal("Error loading .env.development file")
    }

	logglyToken = os.Getenv("LOGGLY_TOKEN")
    host = os.Getenv("HOST")
    clientPort = os.Getenv("CLIENT_PORT")
	serverPort = os.Getenv("SERVER_PORT")
    redisPort = os.Getenv("REDIS_PORT")

	redisClient := redis.NewClient(&redis.Options{
	Addr: host + ":" + redisPort,
    Password: "",
    DB: 0,
    })
    defer redisClient.Close()

	redisdb.InitializeClient(redisClient)
	router := gin.Default()
	router.Use(rateLimitMiddleware())
	router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://" + host + ":" + clientPort},
    AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
	AllowHeaders:     []string{"Content-Type"},
  }))

	v1 := router.Group("/api/v1")
	{
		v1.POST("/scrape", func(c *gin.Context) {
			logger.HandleAPIEndpoint(c, l, h, host, scrape.DisplayScrapingResults)
		})
		v1.POST("/results", func(c *gin.Context) {
			logger.HandleAPIEndpoint(c, l, h, host, scrape.SaveScrapingResults)
		})
		v1.GET("/results", func(c *gin.Context) {
			logger.HandleAPIEndpoint(c, l, h, host, scrape.GetScrapingResults)
		})
		v1.DELETE("/results/:id", func(c *gin.Context) {
			logger.HandleAPIEndpoint(c, l, h, host, scrape.DeleteScrapingResults)
		})
	}
	
	if err := router.Run(":" + serverPort); err != nil {
		l.Fatal("Error starting server: ", err)
	}
}