package dynamodb

import (
	"strings"

	config "server/internal"
	"server/pkg/models"

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

func SaveResults(link string, keywords []string, resultName string, sentences []string, resultID, timestamp string) error {
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

	searchText := strings.Join(sentences, " ")

    item := models.ResultItem{
        ResultID: resultID,
		ResultName: resultName,
        Text:     []string{searchText},
        Keywords: keywords,
        Link:     link,
        Timestamp: timestamp,
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

func GetResults() ([]ResultItem, error) {
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

func DeleteResult(resultId string) error {
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

    input := &dynamodb.DeleteItemInput{
        TableName: aws.String("results"),
        Key: map[string]*dynamodb.AttributeValue{
            "result_id": {
                S: aws.String(resultId),
            },
        },
    }

    _, err = svc.DeleteItem(input)
    if err != nil {
        return err
    }

    return nil
}