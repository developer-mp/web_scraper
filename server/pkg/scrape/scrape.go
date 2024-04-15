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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data: " + err.Error()})
		return
	}

	sentences, err := ScrapeWebpage(form.URL, form.Keywords)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate scraping results: " + err.Error()})
		return
	}

	if len(sentences) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No sentences found for the given keywords"})
		return
	}

	c.JSON(http.StatusOK, sentences)
}

func ScrapeWebpage(url string, keywords []string) ([]string, error) {
	var sentences []string

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
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
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data: " + err.Error()})
        return
    }

	resultId := hash.GenerateHashId()
    currentTime := time.Now().Format(time.RFC3339)

    if err := SaveResultsToRedis(requestData.URL, requestData.Keywords, requestData.ResultName, requestData.Sentences, resultId, currentTime); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save results to Redis: " + err.Error()})
        return
    }

    err := dynamodb.SaveResults(requestData.URL, requestData.Keywords, requestData.ResultName, requestData.Sentences, resultId, currentTime)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save results to database: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Scraping results saved successfully"})
}

func GetScrapingResults(c *gin.Context) {
    cachedResults, err := getResultsFromRedis()
    if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve results from Redis: " + err.Error()})
        return
    }

    if cachedResults != nil {
        c.JSON(http.StatusOK, cachedResults)
        return
    }
	
	results, err := dynamodb.GetResults()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve results from database: " + err.Error()})
		return
	}

	if len(results) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No results found"})
		return
	}

	c.JSON(http.StatusOK, results)
}

func DeleteScrapingResults(c *gin.Context) {
	id := c.Param("id")

	if err := dynamodb.DeleteResult(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete result from database: " + err.Error()})
		return
	}

	if err := deleteResultFromRedis(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete result from Redis: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Result deleted successfully"})
}

func SaveResultsToRedis(url string, keywords []string, resultName string, sentences []string, resultID, timestamp string) error {
    combinedText := strings.Join(sentences, " ")
	
	data := map[string]interface{}{
		"result_id": resultID,
		"result_name": resultName,
		"text": combinedText,
		"keywords": keywords,
        "link": url,      
        "timestamp": timestamp,
    }

    jsonData, err := json.Marshal(data)
    if err != nil {
        return err
    }

    key := resultID
    _, err = redisClient.Set(context.Background(), key, jsonData, 24*time.Hour).Result()
    if err != nil {
    	log.Printf("Failed to save result to Redis: %v", err)
    	return err
	}
	return nil
}

func getResultsFromRedis() (interface{}, error) {
	var results []interface{}
    keys, err := redisClient.Keys(context.Background(), "*").Result()
    if err != nil {
        return nil, err
    }

    for _, key := range keys {
        jsonData, err := redisClient.Get(context.Background(), key).Result()
        if err != nil {
            continue
        }

        var result interface{}
        if err := json.Unmarshal([]byte(jsonData), &result); err != nil {
            continue
        }
        results = append(results, result)
    }

    if len(results) == 0 {
        return nil, nil
    }

    return results, nil
}

func deleteResultFromRedis(resultID string) error {
	_, err := redisClient.Del(context.Background(), resultID).Result()
	if err != nil {
		return err
	}
	return nil
}