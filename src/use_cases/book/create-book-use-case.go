package book

import (
	"log"

	extensions "github.com/YagoFerreira39/go-bookclub/src/adapters/extensions/book"
	book_repository "github.com/YagoFerreira39/go-bookclub/src/adapters/repositories/mongodb"
	dtos "github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/dtos/book"
	requests "github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/requests/book"
)

type CreateBookUseCase struct {
	CreateBookExtension *extensions.CreateBookExtension
	BookRepository      *book_repository.BookRepository
}

func (createBookUseCase *CreateBookUseCase) CreateBook(request requests.CreateBookRequest) (*dtos.CreateBookDto, error) {
	entity := createBookUseCase.CreateBookExtension.FromRequestToEntity(request)

	model := createBookUseCase.CreateBookExtension.FromEntityToModel(entity)

	// calls BookRepository to create a Book here
	result, err := createBookUseCase.BookRepository.CreateBook(model)

	if err != nil {
		log.Printf("Unable to create book.")
		return &dtos.CreateBookDto{}, err
	}

	dto := createBookUseCase.CreateBookExtension.FromModelToDto(result)

	return &dto, nil
}
