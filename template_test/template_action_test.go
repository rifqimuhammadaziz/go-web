package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type Page struct {
	Title string
	Name  string
}

func TemplateActionIf(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../templates/if.gohtml"))
	t.ExecuteTemplate(rw, "if.gohtml", Page{
		Title: "Template Action If",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}