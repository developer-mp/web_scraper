package main

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type ResultItem struct {
	ResultID string   `json:"result_id"`
	Text     []string `json:"text"`
}

func storeInDynamoDB(sentences []string) error {
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

    resultID := generateHashID()
	searchText := strings.Join(sentences, " ")

    item := ResultItem{
        ResultID: resultID,
        Text:     []string{searchText},
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