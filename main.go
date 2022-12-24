package main

import (
	"fmt"
	mongoConnection "go_sample/src/main/config/mongo"
	configurations_yaml "go_sample/src/main/config/yaml"
	"go_sample/src/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.GenerateRoutes()

	mongoConnection.GetDatabaseBean()

	log.Printf("Escutando na porta %s", configurations_yaml.GetBeanPropertyByName("Application.Port"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", configurations_yaml.GetBeanPropertyByName("Application.Port")), r))
}
