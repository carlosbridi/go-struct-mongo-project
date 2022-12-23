package routes

import "github.com/gorilla/mux"

func GenerateRoutes() *mux.Router {

	r := mux.NewRouter()
	return ExportRoutes(r)
}
