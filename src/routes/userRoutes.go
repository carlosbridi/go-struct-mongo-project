package routes

import (
	"go_sample/src/main/controllers"
	"net/http"
)

var userRoutes = []Routes{

	{
		URI:      "/hello",
		Method:   http.MethodGet,
		Function: controllers.HelloWorld,
	},
	{
		URI:      "/customer",
		Method:   http.MethodPost,
		Function: controllers.AddCustomer,
	},
}
