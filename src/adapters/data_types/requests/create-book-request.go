package requests

type CreateBookRequest struct {
	Name      string  `json:"name" bson:"Name"`
	Author    string  `json:"author" bson:"Author"`
	ISBN      string  `json:"isbn" bson:"ISBN"`
	Published float64 `json:"published" bson:"Published"`
}
