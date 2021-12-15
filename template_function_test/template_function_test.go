package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name Is " + myPage.Name
}

func TemplateFunction(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Xenosty"}}`))
	t.ExecuteTemplate(rw, "FUNCTION", MyPage{
		Name: "Xenosty",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobal(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len "Xenosty"}}`))
	t.ExecuteTemplate(rw, "FUNCTION", MyPage{
		Name: "Test Template Function Global",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobalCreate(rw http.ResponseWriter, r *http.Request) {
	t := template.New("GlobalFunction")

	// create template function
	t = t.Funcs(map[string]interface{}{
		// functionName: createFunction | upper(value string)
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	// execute template function upper(Name)
	t = template.Must(t.Parse(`{{ upper .Name }}`))

	t.ExecuteTemplate(rw, "GlobalFunction", MyPage{
		Name: "Test Template Function Global",
	})
}

func TestTemplateFunctionGlobalCreate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobalCreate(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobalCreatePipeline(rw http.ResponseWriter, r *http.Request) {
	t := template.New("GlobalFunction")

	// create template function
	t = t.Funcs(map[string]interface{}{
		// functionName: sayHello | sayHello(name string)
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		// functionName: createFunction | upper(value string)
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	// execute sayHello, then the result throw to upper parameter
	t = template.Must(t.Parse(`{{ sayHello .Name  | upper }}`))

	t.ExecuteTemplate(rw, "GlobalFunction", MyPage{
		Name: "Xenosty",
	})
}

func TestTemplateFunctionGlobalCreatePipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobalCreatePipeline(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
