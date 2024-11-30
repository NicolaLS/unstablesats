package main

import "github.com/NicolaLS/unstablesats/internal/api"
import "net/http"
import "fmt"

func main() {
	http.HandleFunc("/health", api.Health)
	http.HandleFunc("/blink_webhook", api.Blink_webhook)

	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
