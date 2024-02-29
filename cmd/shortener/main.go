package main

import (
	"fmt"
	"net/http"

	api "github.com/krevetkou/yandex-project/internal/api"
)

func main() {
	shortener := api.NewURLShortener()

	http.HandleFunc("/POST", shortener.HandleShorten)
	http.HandleFunc("/GET/", shortener.HandleRedirect)

	fmt.Println("URL Shortener is running on :8080")
	http.ListenAndServe(":8080", nil)
}
