package routes

import (
	"github.com/YagoFerreira39/go-bookclub/src/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookClubRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
}
