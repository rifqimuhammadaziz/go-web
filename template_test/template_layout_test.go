package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"../templates/layout/header.gohtml",
		"../templates/layout/body.gohtml",
		"../templates/layout/footer.gohtml",
		"../templates/layout/layout.gohtml",
	))

	t.ExecuteTemplate(rw, "layout", map[string]interface{}{
		"Title": "Test Template Layout",
		"Name":  "Xenosty",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
