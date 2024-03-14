package dtos

type CreateBookDto struct {
	Id_       string  `json:"Id_"`
	Name      string  `json:"Name"`
	Author    string  `json:"Author"`
	ISBN      string  `json:"ISBN"`
	Published float64 `json:"Published"`
}
