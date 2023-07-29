package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", ping).Methods("GET")

	http.ListenAndServe(":8000", r)
}

func ping(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
