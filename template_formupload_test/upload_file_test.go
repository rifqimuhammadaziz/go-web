package goweb

import (
	"bytes"
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

//go:embed templates/*.gohtml
var templates embed.FS

// parsing one time (global var), faster than inside handler
var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func UploadForm(rw http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(rw, "upload.form.gohtml", nil)
}

func Upload(rw http.ResponseWriter, r *http.Request) {
	// get file
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	// create destination file
	fileDestination, err := os.Create("../resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	// save file to destination
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	// get input text
	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(rw, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm) //
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("../resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// UNIT TESTING FOR UPLOAD FILE & TEXT

//go:embed resources/rifqi.jpg
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Rifqi Muhammad Aziz")
	file, _ := writer.CreateFormFile("file", "TestUpload.jpg")
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
