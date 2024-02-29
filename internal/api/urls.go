package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	services "github.com/krevetkou/yandex-project/internal/services"
)

type URLShortener struct {
	urls map[string]string
}

func NewURLShortener() *URLShortener {
	return &URLShortener{
		urls: make(map[string]string),
	}
}

func (us *URLShortener) HandleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	originalURL, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}

	if string(originalURL) == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	// Generate a unique shortened key for the original URL
	shortKey := services.GenerateShortKey()
	us.urls[shortKey] = string(originalURL)

	// Construct the full shortened URL
	shortenedURL := fmt.Sprintf("http://localhost:8080/GET/%s", shortKey)

	// Render the HTML response with the shortened URL
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, shortenedURL)
}

func (us *URLShortener) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	shortKey := strings.TrimPrefix(r.URL.Path, "/GET/")
	if shortKey == "" {
		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
		return
	}

	// Retrieve the original URL from the `urls` map using the shortened key
	originalURL, found := us.urls[shortKey]
	if !found {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	fmt.Println(originalURL)

	w.Header().Set("Location", originalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
