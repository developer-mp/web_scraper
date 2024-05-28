package main

import (
	"fmt"
	"os"
	logger "server/internal/logger"
	proxy "server/internal/proxy"
	ratelimiter "server/internal/ratelimiter"
	"server/pkg/redisdb"
	"server/pkg/scrape"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/sebest/logrusly"
	"github.com/sirupsen/logrus"
)

var logglyToken string
var host string
var clientPort string
var serverPort string
var redisPort string

func main() {
	err := proxy.InitProxy("appconfig.json")
    if err != nil {
        fmt.Println("Error initializing proxy: ", err)
        return
    }
	
	l := logrus.New()
	h := logrusly.NewLogglyHook(logglyToken, host, logrus.WarnLevel, "backend")
	l.Hooks.Add(h)

    err = godotenv.Load(".env.development")
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
	rateLimiter := ratelimiter.NewRateLimiter(10, 5)
	router.Use(rateLimiter.Limit())
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