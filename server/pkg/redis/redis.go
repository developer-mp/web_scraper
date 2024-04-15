package redisdb

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func InitializeClient(newClient *redis.Client) {
    client = newClient
}

func SaveResults(key string, data map[string]interface{}, expiration time.Duration) error {
    jsonData, err := json.Marshal(data)
    if err != nil {
        return err
    }

    _, err = client.Set(context.Background(), key, jsonData, expiration).Result()
    if err != nil {
        log.Printf("Failed to save result to Redis: %v", err)
        return err
    }
    return nil
}

func GetResults() (interface{}, error) {
    var results []interface{}
    keys, err := client.Keys(context.Background(), "*").Result()
    if err != nil {
        return nil, err
    }

    for _, key := range keys {
        jsonData, err := client.Get(context.Background(), key).Result()
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

func DeleteResult(key string) error {
    _, err := client.Del(context.Background(), key).Result()
    if err != nil {
        return err
    }
    return nil
}
