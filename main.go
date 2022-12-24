package main

import (
	"fmt"
	"go_sample/src/routes"
	"log"
	"net/http"

	configurations_yaml "go_sample/src/main/config/yaml"
)

func main() {
	r := routes.GenerateRoutes()

	log.Printf("Escutando na porta %s", configurations_yaml.GetBeanPropertyByName("Application.Port"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", configurations_yaml.GetBeanPropertyByName("Application.Port")), r))
}
