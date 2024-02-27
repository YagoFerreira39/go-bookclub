package httpconfig

import (
	"log"
	"net/http"

	"github.com/YagoFerreira39/go-bookclub/src/externals/routes"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterHttpConfig() {
	root_path := "/"
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	routes.RegisterBookClubRoutes(router)

	http.Handle(root_path, router)
	log.Fatal(http.ListenAndServe("localhost:8001", router))
}
