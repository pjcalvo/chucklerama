package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
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

func notFound(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Page not Found")
}

type NewJoke struct {
	Joke string `json:"joke"`
}

func handleNewJoke(w http.ResponseWriter, r *http.Request) {
	var joke NewJoke
	err := json.NewEncoder(w).Encode(&joke)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error encoding joke %w", err)
	}
	slog.Info("New Joke", slog.Any("joke", joke))
	w.WriteHeader(http.StatusCreated)
}

func main() {
	handler := http.NewServeMux()
	// html routes
	handler.HandleFunc("GET /", serveHtml(resources.HomePage))
	handler.HandleFunc("GET /new", serveHtml(resources.NewJokePage))
	handler.HandleFunc("GET /terms", serveHtml(resources.TermsPage))
	handler.Handle("", http.NotFoundHandler())

	// api routes
	http.HandleFunc("POST /jokes", handleNewJoke)

	fmt.Println("Starting server on :8080")

	log.Fatal(http.ListenAndServe(":8080", handler))
}
