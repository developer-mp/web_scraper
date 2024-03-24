package dynamodb

import (
	"strings"
	"time"

	config "server/internal"
	hash "server/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type ResultItem struct {
	ResultID string   `json:"result_id"`
	Text     []string `json:"text"`
    Keywords []string `json:"keywords"`
    Link     string   `json:"link"`
    Timestamp string   `json:"timestamp"`
}

func StoreInDynamoDB(sentences []string, keywords []string, link string) error {
    AWSConfig, err := config.ReadAppConfig("appconfig.json")
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

    resultID := hash.GenerateHashID()
	searchText := strings.Join(sentences, " ")
    currentTime := time.Now().Format(time.RFC3339)

    item := ResultItem{
        ResultID: resultID,
        Text:     []string{searchText},
        Keywords: keywords,
        Link:     link,
        Timestamp: currentTime,
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

    return nil
}