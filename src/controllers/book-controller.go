package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/YagoFerreira39/go-bookclub/src/adapters/data_types/requests"
)

func CreateBook(responseWriter http.ResponseWriter, request *http.Request) {
	var createBookRequest requests.CreateBookRequest

	error := json.NewDecoder(request.Body).Decode(&createBookRequest)
	if error != nil {
		fmt.Println(error)
	}

}
