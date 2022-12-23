package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	URI      string
	Method   string
	Function func(w http.ResponseWriter, r *http.Request)
}

func ExportRoutes(r *mux.Router) *mux.Router {
	rotas := userRoutes

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Function).Methods(rota.Method)
	}

	return r
}
