package routes

import (
	book_repository_dynamodb "github.com/YagoFerreira39/go-bookclub/src/adapters/repositories/dynamodb"
	"github.com/YagoFerreira39/go-bookclub/src/controllers"
	awsdynamodb "github.com/YagoFerreira39/go-bookclub/src/externals/infrastructure/aws-dynamodb"
	"github.com/gorilla/mux"
)

var dynamoDBInfrastructure = awsdynamodb.NewDynamoDBInfrastructure()
var bookRepository = book_repository_dynamodb.BookRepository{DynamoDBInfrastructure: dynamoDBInfrastructure}
var bookController = controllers.BookController{BookRepository: &bookRepository}

var RegisterBookClubRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", bookController.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", bookController.GetBookById).Methods("GET")
}
