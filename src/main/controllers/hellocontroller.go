package controllers

import (
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Bridi"))

	//responses.JSON(w, http.StatusCreated, []byte("Bridi"))
}
