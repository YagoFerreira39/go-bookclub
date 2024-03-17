package book

import (
	"log"

	extensions "github.com/YagoFerreira39/go-bookclub/src/adapters/extensions/book"
	book_repository "github.com/YagoFerreira39/go-bookclub/src/adapters/repositories/mongodb"
	"github.com/YagoFerreira39/go-bookclub/src/domain/models"
	dtos "github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/dtos/book"
	requests "github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/requests/book"
)

type GetBookByIdUseCase struct {
	GetBookByIdExtension *extensions.GetBookByIdExtension
	BookRepository       *book_repository.BookRepository
}

func (getBookByIdUseCase *GetBookByIdUseCase) GetBookById(request *requests.GetBookByIdRequest) (*dtos.BookDto, error) {
	bookModelFromDatabase, error := getBookByIdUseCase.getBookById(request.BookId)

	dto := getBookByIdUseCase.createBookDto(bookModelFromDatabase)

	return dto, error
}

func (getBookByIdUseCase *GetBookByIdUseCase) getBookById(bookId string) (*models.BookModel, error) {
	bookModelFromDatabase, error := getBookByIdUseCase.BookRepository.GetBookById(bookId)

	if error != nil {
		log.Printf("Unable to find book.")
	}

	return &bookModelFromDatabase, nil
}

func (getBookByIdUseCase *GetBookByIdUseCase) createBookDto(bookModelFromDatabase *models.BookModel) *dtos.BookDto {
	dto := getBookByIdUseCase.GetBookByIdExtension.FromModelToDto(bookModelFromDatabase)

	return &dto
}
