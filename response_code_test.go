package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		rw.WriteHeader(http.StatusBadRequest) // rw.WriteHeader(400)
		fmt.Fprint(rw, "Name is empty!")
	} else {
		fmt.Fprintf(rw, "Hello %s", name) // default status success (200) OK
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	// test invalid (without request header 'name')
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode) // 400
	fmt.Println(response.Status)     // 400 Bad Request
	fmt.Println(string(body))        // print 'Name is empty!'
}

func TestResponseCodeValid(t *testing.T) {
	// test valid (with request header 'name')
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=Xenosty", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode) // 200
	fmt.Println(response.Status)     // 200 OK
	fmt.Println(string(body))        // print 'Hello Xenosty'
}
