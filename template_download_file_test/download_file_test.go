package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(rw http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	if file == "" {
		rw.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(rw, "Bad Request")
		return // stop function if no file
	}

	// auto download file (without show file) if has request to file
	rw.Header().Add("Content-Disposition", "attactment; filename=\""+file+"\"")
	http.ServeFile(rw, r, "../resources/"+file)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
