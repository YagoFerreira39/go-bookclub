package routes

import (
	book_repository "github.com/YagoFerreira39/go-bookclub/src/adapters/repositories/mongodb"
	"github.com/YagoFerreira39/go-bookclub/src/controllers"
	"github.com/YagoFerreira39/go-bookclub/src/externals/infrastructure/mongodb"
	"github.com/gorilla/mux"
)

var mongoDBInfrastructure, _ = mongodb.NewMongoDBInfrastructure()
var bookRepository = book_repository.BookRepository{MongoDBInfrastructure: mongoDBInfrastructure}
var bookController = controllers.BookController{BookRepository: &bookRepository}

var RegisterBookClubRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", bookController.CreateBook).Methods("POST")
	router.HandleFunc("/book/", bookController.GetBookById).Methods("GET")
}
