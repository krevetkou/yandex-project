package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	port = "8080"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/update/", handleRequest)
	log.Println(fmt.Sprintf("Starting web server http://127.0.0.1:%s", port))
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	log.Fatal(err)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	url := r.URL.Path

	splitURL := strings.Split(url, "/")
	fmt.Println(splitURL)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "text/plain")
}
