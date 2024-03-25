package book

import (
	"log"

	extensions "github.com/YagoFerreira39/go-bookclub/src/adapters/extensions/book"
	book_repository_dynamodb "github.com/YagoFerreira39/go-bookclub/src/adapters/repositories/dynamodb"
	"github.com/YagoFerreira39/go-bookclub/src/domain/entities"
	"github.com/YagoFerreira39/go-bookclub/src/domain/models"
	dtos "github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/dtos/book"
	requests "github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/requests/book"
)

type CreateBookUseCase struct {
	CreateBookExtension *extensions.CreateBookExtension
	BookRepository      *book_repository_dynamodb.BookRepository
}

func (createBookUseCase *CreateBookUseCase) CreateBook(request *requests.CreateBookRequest) (*dtos.CreateBookDto, error) {
	entity := createBookUseCase.createBookEntity(request)

	model := createBookUseCase.createBookModel(entity)

	result, error := createBookUseCase.createBook(model)

	dto := createBookUseCase.CreateBookExtension.FromModelToDto(result)

	return dto, error
}

// Private methods
func (createBookUseCase *CreateBookUseCase) createBookEntity(request *requests.CreateBookRequest) *entities.BookEntity {
	entity := createBookUseCase.CreateBookExtension.FromRequestToEntity(request)

	return entity
}

func (createBookUseCase *CreateBookUseCase) createBookModel(entity *entities.BookEntity) *models.BookModel {
	model := createBookUseCase.CreateBookExtension.FromEntityToModel(entity)

	return model
}

func (createBookUseCase *CreateBookUseCase) createBook(model *models.BookModel) (*models.BookModel, error) {
	result, error := createBookUseCase.BookRepository.CreateBook(model)

	if error != nil {
		log.Printf("Unable to create book.")
		return &models.BookModel{}, error
	}

	return result, nil
}
