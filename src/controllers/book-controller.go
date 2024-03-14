package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	extensions "github.com/YagoFerreira39/go-bookclub/src/adapters/extensions/book"
	book_repository "github.com/YagoFerreira39/go-bookclub/src/adapters/repositories/mongodb"
	"github.com/YagoFerreira39/go-bookclub/src/use_cases/book"
)

type BookController struct {
	BookRepository *book_repository.BookRepository
}

func (controller *BookController) CreateBook(responseWriter http.ResponseWriter, request *http.Request) {
	createBookExtension := &extensions.CreateBookExtension{}
	createBookUseCase := book.CreateBookUseCase{
		CreateBookExtension: createBookExtension,
		BookRepository:      controller.BookRepository,
	}

	createBookRequest, error := createBookExtension.FromRouterRequestToRequest(request)
	if error != nil {
		fmt.Println(error)
	}

	useCaseResponse, error := createBookUseCase.CreateBook(createBookRequest)
	if error != nil {
		http.Error(responseWriter, "Error creating book", http.StatusInternalServerError)
		fmt.Println("Error creating book:", error)
		return
	}

	response, _ := json.Marshal(useCaseResponse)

	// Handle successful book creation
	responseWriter.Header().Set("Content-Type", "pkglication/json")
	responseWriter.WriteHeader(http.StatusCreated)
	responseWriter.Write(response)
}
