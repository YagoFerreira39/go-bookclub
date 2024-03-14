package responses

type BaseApiResponse struct {
	Status    bool   `json:"status"`
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}