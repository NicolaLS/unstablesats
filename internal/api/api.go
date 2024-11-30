package api

import "github.com/NicolaLS/unstablesats/internal/models"

import "net/http"
import "fmt"
import "encoding/json"

func Health(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "OK")
}

func Blink_webhook(w http.ResponseWriter, r *http.Request) {
	var payload models.BlinkWebhookPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
	}

	fmt.Printf("Received webhook payload: %+v\n", payload)
	w.WriteHeader(http.StatusOK)
}
