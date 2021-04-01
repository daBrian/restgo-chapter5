package main

import (
	"chapter5/domain"
	"chapter5/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func createInvoiceHandler(writer http.ResponseWriter, request http.Request) {
	// Read invoice data from request body
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	// CreateInvoice invoice and unmarshal it from JSON
	var i domain.Invoice
	json.Unmarshal(body, &i)

	i.CustomerId, _ = strconv.Atoi(mux.Vars(request)["customerId"])
	created := repository.CreateInvoice(i)
	b, err := json.Marshal(created)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Write response
	location := fmt.Sprintf("%s/%d", request.URL.String(), created.Id)
	writer.Header().Set("Location", location)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(b)
}
