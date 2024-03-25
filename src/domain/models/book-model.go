package models

type BookModel struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Author    string  `json:"author"`
	ISBN      string  `json:"isbn"`
	Published float64 `json:"published"`
}