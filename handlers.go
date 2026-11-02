package main

import (
	"bytes"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFS(siteFiles, "templates/index.html", "templates/header.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/#about", http.StatusMovedPermanently)
}

func cvHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	pdf, err := fs.ReadFile(siteFiles, "Ammar_CV.pdf")
	if err != nil {
		log.Println(err)
		http.Error(w, "CV unavailable", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	http.ServeContent(w, r, "Ammar_CV.pdf", time.Time{}, bytes.NewReader(pdf))
}
