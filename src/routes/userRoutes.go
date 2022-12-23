package routes

import (
	"main/src/main/controllers"
	"net/http"
)

var userRoutes = []Routes{

	{
		URI:      "/hello",
		Method:   http.MethodGet,
		Function: controllers.HelloWorld,
	},
}
