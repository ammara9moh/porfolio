package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed templates static Ammar_CV.pdf
var siteFiles embed.FS

func main() {
	staticFiles, err := fs.Sub(siteFiles, "static")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFiles))))
	http.HandleFunc("/Ammar_CV.pdf", cvHandler)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	address := ":" + port
	log.Printf("Listening on http://localhost%s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
