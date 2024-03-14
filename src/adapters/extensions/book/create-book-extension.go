package extensions

import (
	"encoding/json"
	"net/http"

	"github.com/YagoFerreira39/go-bookclub/src/domain/entities"
	"github.com/YagoFerreira39/go-bookclub/src/domain/models"
	dtos "github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/dtos/book"
	requests "github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/requests/book"
)

type CreateBookExtension struct {
}

func (createBookExtension *CreateBookExtension) FromRouterRequestToRequest(request *http.Request) (requests.CreateBookRequest, error) {
	var createBookRequest requests.CreateBookRequest
	error := json.NewDecoder(request.Body).Decode(&createBookRequest)

	return createBookRequest, error
}

func (createBookExtension *CreateBookExtension) FromRequestToEntity(request requests.CreateBookRequest) entities.BookEntity {
	entity := entities.BookEntity{
		Name:      request.Name,
		Author:    request.Author,
		ISBN:      request.ISBN,
		Published: request.Published,
	}

	return entity
}

func (createBookExtension *CreateBookExtension) FromEntityToModel(entity entities.BookEntity) models.BookModel {
	model := models.BookModel{
		Name:      entity.Name,
		Author:    entity.Author,
		ISBN:      entity.ISBN,
		Published: entity.Published,
	}

	return model
}

func (createBookExtension *CreateBookExtension) FromModelToDto(model models.BookModel) dtos.CreateBookDto {
	dto := dtos.CreateBookDto{
		Id_:       model.ID.Hex(),
		Name:      model.Name,
		Author:    model.Author,
		ISBN:      model.ISBN,
		Published: model.Published,
	}

	return dto
}
