package controllers

import (
	"encoding/json"
	"fmt"
	"go_sample/src/main/domains"
	"go_sample/src/main/responses"
	"io/ioutil"
	"net/http"
)

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var customer domains.Customer
	if erro = json.Unmarshal(corpoRequest, &customer); erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	fmt.Println(customer)

}
