package scrape

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	dynamodb "server/pkg/db"
	hash "server/utils"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitializeRedisClient(client *redis.Client) {
    redisClient = client
}


func DisplayScrapingResults(c *gin.Context) {
	var form struct {
		URL      string   `form:"url"`
		Keywords []string `form:"keywords"`
	}

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sentences, err := ScrapeWebpage(form.URL, form.Keywords)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(sentences) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No sentences found for the given keywords."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": sentences})
}

func ScrapeWebpage(url string, keywords []string) ([]string, error) {
	var sentences []string

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		for _, keyword := range keywords {
			if strings.Contains(text, keyword) {
				sentences = append(sentences, text)
				break
			}
		}
	})

	return sentences, nil
}

func SaveScrapingResults(c *gin.Context) {
	if redisClient == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis client is not initialized"})
        return
    }

    var requestData struct {
        URL        string   `json:"url"`
        Keywords   []string `json:"keywords"`
        ResultName string   `json:"resultName"`
        Sentences  []string `json:"sentences"`
    }

    if err := c.BindJSON(&requestData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	resultID := hash.GenerateHashID()
    currentTime := time.Now().Format(time.RFC3339)

    if err := SaveToRedis(requestData.URL, requestData.Keywords, requestData.ResultName, requestData.Sentences, resultID, currentTime); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save result to Redis"})
        return
    }

    err := dynamodb.SaveResults(requestData.URL, requestData.Keywords, requestData.ResultName, requestData.Sentences, resultID, currentTime)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Scraping results saved successfully"})
}

func GetScrapingResults(c *gin.Context) {
	results, err := dynamodb.GetResults()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

func DeleteScrapingResult(c *gin.Context) {
	id := c.Param("id")

	err := dynamodb.DeleteResult(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete result"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Result deleted successfully"})
}

func SaveToRedis(url string, keywords []string, resultName string, sentences []string, resultID, timestamp string) error {
    data := map[string]interface{}{
        "url":       url,
        "keywords":  keywords,
        "resultName": resultName,
        "sentences": sentences,
		"resultID": resultID,
        "timestamp": timestamp,
    }

    jsonData, err := json.Marshal(data)
    if err != nil {
        return err
    }

    key := resultName
    _, err = redisClient.Set(context.Background(), key, jsonData, 0).Result()
    if err != nil {
    	log.Printf("Failed to save to Redis: %v", err)
    	return err
	}
	return nil
}