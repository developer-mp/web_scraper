package dynamodb

import (
	"strings"

	awsconfig "server/internal/config"
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

func CheckForDuplicate(url string, keywords []string) (bool, error) {
    AWSConfig, err := awsconfig.ReadAWSConfig("appconfig.json")
    if err != nil {
        return false, err
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

    var dbKeywords []*dynamodb.AttributeValue
    for _, keyword := range keywords {
        dbKeywords = append(dbKeywords, &dynamodb.AttributeValue{S: aws.String(keyword)})
    }

    params := &dynamodb.ScanInput{
        TableName:        aws.String("results"),
        FilterExpression: aws.String("#link = :url AND #keywords = :keywords"),
        ExpressionAttributeNames: map[string]*string{
            "#link":     aws.String("link"),
            "#keywords": aws.String("keywords"),
        },
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":url": {
                S: aws.String(url),
            },
            ":keywords": {
                L: dbKeywords,
            },
        },
    }

    result, err := svc.Scan(params)
    if err != nil {
        return false, err
    }

    if len(result.Items) > 0 {
        return true, nil
    }

    return false, nil
}

func SaveResults(link string, keywords []string, resultName string, sentences []string, resultID, timestamp string) error {
    AWSConfig, err := awsconfig.ReadAWSConfig("appconfig.json")
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
	AWSConfig, err := awsconfig.ReadAWSConfig("appconfig.json")
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
    AWSConfig, err := awsconfig.ReadAWSConfig("appconfig.json")
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