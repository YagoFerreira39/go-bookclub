package book_repository_dynamodb

import (
	"fmt"
	"log"

	"github.com/YagoFerreira39/go-bookclub/src/domain/models"
	awsdynamodb "github.com/YagoFerreira39/go-bookclub/src/externals/infrastructure/aws-dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

type BookRepository struct {
	DynamoDBInfrastructure *awsdynamodb.DynamoDBInfrastructure
}

func (repository *BookRepository) CreateBook(model *models.BookModel) (*models.BookModel, error) {
	client := repository.DynamoDBInfrastructure.Client

	model.ID = uuid.New().String()

	attributeValue, error := dynamodbattribute.MarshalMap(model)
	if error != nil {
		log.Fatalln(error)
	}

	toInsert := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String("Books"),
	}

	_, error = client.PutItem(toInsert)
	if error != nil {
		log.Fatalln(error)
		return &models.BookModel{}, error
	}

	log.Println("Book created successfully.")
	return model, nil
}

func (repository *BookRepository) GetBookById(bookId string) (*models.BookModel, error) {
	client := repository.DynamoDBInfrastructure.Client

	keyCondition := &dynamodb.QueryInput{
		TableName: aws.String("Books"),
		KeyConditions: map[string]*dynamodb.Condition{
			"id": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(bookId),
					},
				},
			},
		},
	}

	result, error := client.Query(keyCondition)
	if error != nil {
		fmt.Println("Error:", error)
		return &models.BookModel{}, error
	}

	var book models.BookModel
	if len(result.Items) > 0 {
		err := dynamodbattribute.UnmarshalMap(result.Items[0], &book)
		if err != nil {
			fmt.Println("Error:", err)
			return &models.BookModel{}, err
		}

		fmt.Println("Retrieved Book:", book)

		return &book, nil
	} else {

		fmt.Println("No matching item found")
		return &models.BookModel{}, error
	}

}
