package handlers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "error parsing form: "+err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "error reading file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "error reading file content: "+err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := service.AutoConvert(string(data))
	if err != nil {
		http.Error(w, "conversion error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	newFileName := time.Now().UTC().Format("20060102_150405") + ext

	if err := os.WriteFile(newFileName, []byte(result), 0644); err != nil {
		http.Error(w, "error writing file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Converted result:\n%s\n\nSaved to file: %s", result, newFileName)
}
