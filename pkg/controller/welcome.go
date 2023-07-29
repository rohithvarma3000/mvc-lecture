package controller

import (
	"net/http"

	"github.com/rohithvarma3000/mvc-lecture/pkg/views"
)

func Welcome(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	t := views.StartPage()
	t.Execute(writer, nil)
}
