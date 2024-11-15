package main

import (
	"fmt"
	"net/http"

	"github.com/pjcalvo/chucklerama/internal/resources"
)

func serveHTML(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, resources.HomePage)
}

func main() {
	http.HandleFunc("/", serveHTML)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
