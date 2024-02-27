package requests

type UpdateBookRequest struct {
	Id        string
	Name      string
	Author    string
	ISBN      string
	Published float64
}
