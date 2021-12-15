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

func TemplateActionOperator(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../templates/comparator.gohtml"))
	t.ExecuteTemplate(rw, "comparator.gohtml", map[string]interface{}{
		"Title":      "Test Template Action Operator",
		"FinalValue": 50,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../templates/range.gohtml"))
	t.ExecuteTemplate(rw, "range.gohtml", map[string]interface{}{
		"Title": "Test Template Action Range",
		"Hobbies": []string{
			"Play Dota2", "Learning Go & Cloud", "Watching Serial",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionWith(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../templates/with.gohtml"))
	t.ExecuteTemplate(rw, "with.gohtml", map[string]interface{}{
		"Title": "Test Template With",
		"Name":  "Xenosty",
		"Address": map[string]interface{}{
			"Street":   "Jalan Citarum",
			"City":     "Tegal",
			"Province": "Central Java",
			"Region":   "Indonesia",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
