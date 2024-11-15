package main

import (
	"fmt"
	"net/http"

	"github.com/pjcalvo/chucklerama/internal/resources"
)

func serveHtml(resource string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, resource)
	}
}

func main() {
	http.HandleFunc("/", serveHtml(resources.HomePage))
	http.HandleFunc("/new", serveHtml(resources.NewJoke))

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
