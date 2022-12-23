package main

import (
	"fmt"
	"log"
	"main/src/main/config"
	"main/src/routes"
	"net/http"
)

func main() {
	r := routes.GenerateRoutes()
	fmt.Println("Hello!!")

	log.Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
