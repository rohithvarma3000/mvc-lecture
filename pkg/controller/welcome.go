package controller

import (
	"net/http"

	"github.com/rohithvarma3000/mvc-lecture/pkg/views"
)

func Welcome(writer http.ResponseWriter, request *http.Request) {
	t := views.StartPage()
	t.Execute(writer, nil)
}
