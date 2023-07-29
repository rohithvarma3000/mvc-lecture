package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rohithvarma3000/mvc-lecture/pkg/controller"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Welcome).Methods("GET")
	r.HandleFunc("/add", controller.Add).Methods("POST")
	r.HandleFunc("/list", controller.List).Methods("GET")

	http.ListenAndServe(":8000", r)
}
