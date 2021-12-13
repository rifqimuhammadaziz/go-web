package goweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" { // if address has query parameter 'name', serve ok.html
		http.ServeFile(rw, r, "./resources/ok.html")
	} else {
		http.ServeFile(rw, r, "./resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourceOk string

//go:embed resources/notfound.html
var resourceNotfound string

func ServeFileEmbed(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" { // if address has query parameter 'name', serve ok.html
		fmt.Fprint(rw, resourceOk)
	} else {
		fmt.Fprint(rw, resourceNotfound)
	}
}

func TestServeFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
