package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(rw http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type") // non case sensitive
	fmt.Fprint(rw, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json") // non case sensitive

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func ResponseHeader(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("X-Powered-By", "Xenosty")
	fmt.Fprint(rw, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json") // non case sensitive

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

	fmt.Println(response.Header.Get("x-powered-by")) // get header 'x-powered-by' : Xenosty
}
