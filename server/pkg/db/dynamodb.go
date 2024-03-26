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
	ResultName string `json:"result_name"`
	Text     []string `json:"text"`
    Keywords []string `json:"keywords"`
    Link     string   `json:"link"`
    Timestamp string   `json:"timestamp"`
}

func SaveScrapingResults(link string, keywords []string, resultName string, sentences []string) error {
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
		ResultName: resultName,
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
        TableName: aws.String("results"),
        Item:      itemBytes,
    })
    if err != nil {
        return err
    }

    return nil
}

func GetScrapingResults() ([]ResultItem, error) {
	AWSConfig, err := config.ReadAppConfig("appconfig.json")
	if err != nil {
		return nil, err
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

	params := &dynamodb.ScanInput{
		TableName: aws.String("results"),
	}

	result, err := svc.Scan(params)
	if err != nil {
		return nil, err
	}

	var items []ResultItem
	for _, item := range result.Items {
		var resultItem ResultItem
		err := dynamodbattribute.UnmarshalMap(item, &resultItem)
		if err != nil {
			return nil, err
		}
		items = append(items, resultItem)
	}

	return items, nil
}