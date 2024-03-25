package extensions

import (
	"encoding/json"
	"net/http"

	"github.com/YagoFerreira39/go-bookclub/src/domain/entities"
	"github.com/YagoFerreira39/go-bookclub/src/domain/models"
	dtos "github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/dtos/book"
	requests "github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/requests/book"
	"github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/responses"
	book_responses "github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/responses/book"
	"github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/responses/payload"
)

type CreateBookExtension struct{}

func (createBookExtension *CreateBookExtension) FromRouterRequestToRequest(request *http.Request) (*requests.CreateBookRequest, error) {
	var createBookRequest requests.CreateBookRequest
	error := json.NewDecoder(request.Body).Decode(&createBookRequest)

	return &createBookRequest, error
}

func (createBookExtension *CreateBookExtension) FromRequestToEntity(request *requests.CreateBookRequest) *entities.BookEntity {
	entity := entities.NewBookEntity(nil, request.Name, request.Author, request.ISBN, request.Published)

	return entity
}

func (createBookExtension *CreateBookExtension) FromEntityToModel(entity *entities.BookEntity) *models.BookModel {
	model := models.BookModel{
		Name:      entity.Name,
		Author:    entity.Author,
		ISBN:      entity.ISBN,
		Published: entity.Published,
	}

	return &model
}

func (createBookExtension *CreateBookExtension) FromModelToDto(model *models.BookModel) *dtos.CreateBookDto {
	dto := dtos.CreateBookDto{
		ID:        model.ID,
		Name:      model.Name,
		Author:    model.Author,
		ISBN:      model.ISBN,
		Published: model.Published,
	}

	return &dto
}

func (createBookExtension *CreateBookExtension) FromDtoToResponse(dto *dtos.CreateBookDto) *book_responses.CreateBookResponse {
	response := book_responses.CreateBookResponse{
		BaseApiResponse: responses.BaseApiResponse{
			Status:    true,
			ErrorCode: 0,
			Message:   "Book create successfully.",
		},
		Payload: payload.CreateBookResponsePayload{
			ID:        dto.ID,
			Name:      dto.Name,
			Author:    dto.Author,
			ISBN:      dto.ISBN,
			Published: dto.Published,
		},
	}

	return &response
}
