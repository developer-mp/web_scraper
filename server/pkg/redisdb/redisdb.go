package redisdb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"server/pkg/models"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitializeClient(newClient *redis.Client) {
    if newClient == nil {
        fmt.Println("Redis client is not initialized")
        return
    }
    redisClient = newClient
}

func SaveResults(link string, keywords []string, resultName string, sentences []string, resultId, timestamp string) error {
    combinedText := strings.Join(sentences, " ")

    item := models.ResultItem{
        ResultID:   resultId,
        ResultName: resultName,
        Text:       []string{combinedText},
        Keywords:   keywords,
        Link:       link,
        Timestamp:  timestamp,
    }

    jsonData, err := json.Marshal(item)
    if err != nil {
        return err
    }

    key := resultId
    _, err = redisClient.Set(context.Background(), key, jsonData, 24*time.Hour).Result()
    if err != nil {
    	log.Printf("Failed to save result to Redis: %v", err)
    	return err
	}
	return nil
}

func GetResults() (interface{}, error) {
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

func DeleteResult(resultId string) error {
    _, err := redisClient.Del(context.Background(), resultId).Result()
    if err != nil {
        return err
    }
    return nil
}
