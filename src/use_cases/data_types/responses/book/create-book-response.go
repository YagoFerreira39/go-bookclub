package book_responses

import (
	"github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/responses"
	"github.com/YagoFerreira39/go-bookclub/src/use_cases/data_types/responses/payload"
)

type CreateBookResponse struct {
	responses.BaseApiResponse
	Payload payload.CreateBookResponsePayload `json:"payload"`
}
