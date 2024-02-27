package extensions

import (
	"encoding/json"
	"net/http"

	"github.com/YagoFerreira39/go-bookclub/src/adapters/data_types/requests"
	"github.com/YagoFerreira39/go-bookclub/src/domain/entities"
	"github.com/YagoFerreira39/go-bookclub/src/domain/models"
)

func FromRouterRequestToRequest(request *http.Request) (requests.CreateBookRequest, error) {
	var createBookRequest requests.CreateBookRequest
	error := json.NewDecoder(request.Body).Decode(&createBookRequest)

	return createBookRequest, error
}

func FromRequestToEntity(request requests.CreateBookRequest) entities.BookEntity {
	entity := entities.BookEntity{
		Name:      request.Name,
		Author:    request.Author,
		ISBN:      request.ISBN,
		Published: request.Published,
	}

	return entity
}

func FromEntityToModel(entity entities.BookEntity) models.BookModel {
	model := models.BookModel{
		Name:      entity.Name,
		Author:    entity.Author,
		ISBN:      entity.ISBN,
		Published: entity.Published,
	}

	return model
}
