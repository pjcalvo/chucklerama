package internal

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

func ServeHtml(resource string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Serving HTML", slog.Any("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, resource)
	}
}

func NotFound(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Page not Found")
}

type NewJoke struct {
	Joke string `json:"joke"`
}

func HandleNewJoke(w http.ResponseWriter, r *http.Request) {
	slog.Info("New Joke Request")
	var joke NewJoke
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&joke)
	if err != nil {
		slog.Error("Error encoding joke", slog.String("error", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error encoding joke %w", err)
		return
	}
	slog.Info("New Joke", slog.Any("joke", joke))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Joke Created"))
}
