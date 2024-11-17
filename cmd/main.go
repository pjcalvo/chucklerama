package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pjcalvo/chucklerama/internal"
	"github.com/pjcalvo/chucklerama/resources"
)

func main() {
	handler := http.NewServeMux()
	// html routes
	handler.HandleFunc("GET /", internal.ServeHtml(resources.HomePage))
	handler.HandleFunc("GET /new", internal.ServeHtml(resources.NewJokePage))
	handler.HandleFunc("GET /terms", internal.ServeHtml(resources.TermsPage))
	handler.HandleFunc("GET /thanks", internal.ServeHtml(resources.ThanksPage))

	// api routes
	handler.HandleFunc("POST /jokes", internal.HandleNewJoke)

	fmt.Println("Starting server on :8080")

	log.Fatal(http.ListenAndServe(":8080", handler))
}
