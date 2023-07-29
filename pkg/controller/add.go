package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Add(writer http.ResponseWriter, request *http.Request) {
	var body Book
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	fmt.Println(body.Name)
	writer.WriteHeader(http.StatusOK)
}
