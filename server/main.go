package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"log"
	"net/http"
	"strings"

	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

type AppConfig struct {
	AWS struct {
		AWSAccessKeyID     string `json:"aws_access_key_id"`
		AWSSecretAccessKey string `json:"aws_secret_access_key"`
		AWSRegion          string `json:"aws_region"`
	} `json:"AWS"`
}

type ResultItem struct {
	ResultID string   `json:"result_id"`
	Text     []string `json:"text"`
}

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

func handleScrape(c *gin.Context) {
	var form struct {
		URL      string   `form:"url"`
		Keywords []string `form:"keywords"`
	}

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("URL:", form.URL)
	log.Println("Keywords:", form.Keywords)

	sentences, err := scrapeWebpage(form.URL, form.Keywords)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(sentences) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No sentences found for the given keywords."})
		return
	}

	err = storeInDynamoDB(form.URL, sentences)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": sentences})
}

func scrapeWebpage(url string, keywords []string) ([]string, error) {
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

func storeInDynamoDB(url string, sentences []string) error {
	AWSConfig, err := readAppConfig("appconfig.json")
    if err != nil {
        return err
    }

    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String(AWSConfig.AWS.AWSRegion),
        Credentials: credentials.NewStaticCredentials(
            AWSConfig.AWS.AWSAccessKeyID,
            AWSConfig.AWS.AWSSecretAccessKey,
            "",
        ),
    }))

	svc := dynamodb.New(sess)

	resultID := generateHash(url)

	for _, sentence := range sentences {
        item := ResultItem{
            ResultID: resultID,
            Text:     []string{sentence},
        }

    itemBytes, err := dynamodbattribute.MarshalMap(item)
        if err != nil {
            return err
        }

        _, err = svc.PutItem(&dynamodb.PutItemInput{
            TableName: aws.String("result"),
            Item:      itemBytes,
        })
        if err != nil {
            return err
        }
    }

    return nil
}

func generateHash(data string) string {
    hasher := sha256.New()
    hasher.Write([]byte(data))
    hash := hex.EncodeToString(hasher.Sum(nil))
    return hash
}

func readAppConfig(filename string) (*AppConfig, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var appConfig  AppConfig
    err = json.Unmarshal(data, &appConfig)
    if err != nil {
        return nil, err
    }

    return &appConfig, nil
}