package extensions

import (
	"fmt"
	"net/http"
)

type CreateBookExtension struct{}

func FromRouterRequestToRequest(req *http.Request ) {
	body := req.Body

	fmt.Println(body)

	// request := requests.CreateBookRequest{
	// 	Name: body.Name,
	// 	Author: body.Author,
	// 	ISBN: body.ISBN,
	// 	Published: body.Published,
	// }

}
