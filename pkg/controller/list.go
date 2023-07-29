package controller

import (
	"net/http"

	"github.com/rohithvarma3000/mvc-lecture/pkg/views"
)

func List(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	t := views.ListPage()
	writer.WriteHeader(http.StatusOK)

	test := ListBooks{
		Books: []Book{
			{Name: "B1"},
			{Name: "B2"},
		},
	}
	t.Execute(writer, test)
}
