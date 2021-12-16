package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello Redirect")
}

func RedirectFrom(rw http.ResponseWriter, r *http.Request) {
	http.Redirect(rw, r, "/redirect-to", http.StatusTemporaryRedirect) // redirect to 'localhost/redirect-to'
}

func RedirectOut(rw http.ResponseWriter, r *http.Request) {
	http.Redirect(rw, r, "https://www.google.com", http.StatusTemporaryRedirect) // redirect to google
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom) // if request to '/redirect-to' it will run RedirectFrom function then redirect to '/redirect-to'
	mux.HandleFunc("/redirect-to", RedirectTo)     // if request to '/redirect-to' print 'Hello Redirect'
	mux.HandleFunc("/redirect-out", RedirectOut)   // if request to '/redirect-out' it will run RedirectOut function then redirect to google

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
