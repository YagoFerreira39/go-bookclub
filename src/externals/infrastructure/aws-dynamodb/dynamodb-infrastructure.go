package awsdynamodb

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBInfrastructure struct {
	Client *dynamodb.DynamoDB
}

func NewDynamoDBInfrastructure() *DynamoDBInfrastructure {
	sess, error := session.NewSession(&aws.Config{
		Region:           aws.String("us-west-2"),
		Endpoint:         aws.String("http://localhost:8000"),
		Credentials:      credentials.NewStaticCredentials("placeholder", "placeholder", ""),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	})

	if error != nil {
		log.Panicln(error)
		return nil
	}

	infrastructure := &DynamoDBInfrastructure{
		Client: dynamodb.New(sess),
	}

	return infrastructure
}
