package goweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templates embed.FS

// parsing one time (global var), faster than inside handler
var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func TemplateAutoEscape(rw http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(rw, "post.gohtml", map[string]interface{}{
		"Title": "Test Template Auto Escape",
		"Body":  "<p>This is Body</p><script>You hacked</script>", // by default it will escape or not executed in html/js, output text only
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateAutoEscapeDisabled(rw http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(rw, "post.gohtml", map[string]interface{}{
		"Title": "Test Template Auto Escape",
		"Body":  template.HTML("<p>This is Body</p>"), // rendered in html, auto escape disabled, but potential to xss (cross site scripting)
	})
}

func TestTemplateAutoEscapeDisabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeDisabledServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateXSS(rw http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(rw, "post.gohtml", map[string]interface{}{
		"Title": "Test Template Auto Escape",
		"Body":  template.HTML(r.URL.Query().Get("body")), // rendered in html, auto escape disabled, but potential to xss (cross site scripting)
	})
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?body=<p>Hacked</p>", nil) // XSS, tag html rendered
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
