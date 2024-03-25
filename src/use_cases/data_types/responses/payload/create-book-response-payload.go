package payload

type CreateBookResponsePayload struct {
	ID        string  `json:"ID"`
	Name      string  `json:"Name"`
	Author    string  `json:"Author"`
	ISBN      string  `json:"ISBN"`
	Published float64 `json:"Published"`
}
