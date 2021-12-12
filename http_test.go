package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello World")
}

// for unit testing without running server
func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	bodyString := string(body)
	fmt.Println(bodyString)
}
