package extensions

import (
	"net/http"

	"github.com/YagoFerreira39/go-bookclub/src/domain/models"
	dtos "github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/dtos/book"
	requests "github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/requests/book"
	"github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/responses"
	book_responses "github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/responses/book"
	"github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/responses/payload"
	"github.com/gorilla/mux"
)

type GetBookByIdExtension struct{}

func (getBookByIdExtension *GetBookByIdExtension) FromRouterRequestToRequest(request *http.Request) *requests.GetBookByIdRequest {
	vars := mux.Vars(request)
	bookId := vars["bookId"]

	getBookByIdRequest := requests.GetBookByIdRequest{
		BookId: bookId,
	}

	return &getBookByIdRequest
}

func (getBookByIdExtension *GetBookByIdExtension) FromModelToDto(model *models.BookModel) *dtos.BookDto {
	dto := dtos.BookDto{
		Id_:       model.ID.Hex(),
		Name:      model.Name,
		Author:    model.Author,
		ISBN:      model.ISBN,
		Published: model.Published,
	}

	return &dto
}

func (getBookByIdExtension *GetBookByIdExtension) FromDtoToResponse(dto *dtos.BookDto) *book_responses.GetBookByIdResponse {
	response := book_responses.GetBookByIdResponse{
		BaseApiResponse: responses.BaseApiResponse{
			Status:    true,
			ErrorCode: 0,
			Message:   "Book find with success.",
		},
		Payload: payload.GetBookByIdResponsePayload{
			Id_:       dto.Id_,
			Name:      dto.Name,
			Author:    dto.Author,
			ISBN:      dto.ISBN,
			Published: dto.Published,
		},
	}

	return &response
}
