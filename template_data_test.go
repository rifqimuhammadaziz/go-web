package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	// parsing data
	t.ExecuteTemplate(rw, "name.gohtml", map[string]interface{}{
		"Title": "Testing Template Data Map",
		"Name":  "Xenosty",
		"Address": map[string]interface{}{
			"Street": "Indonesia",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Page struct {
	Title   string
	Name    string
	Address Address
}

type Address struct {
	Street string
}

func TemplateDataStruct(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	// parsing data using 'Page' struct
	t.ExecuteTemplate(rw, "name.gohtml", Page{
		Title: "Testing Template Data Map",
		Name:  "Xenosty",
		Address: Address{ // nested struct
			Street: "Central Java, Indonesia",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
