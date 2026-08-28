package respond

import (
	"encoding/json"
	"log"
	"net/http"
	"server/models"
)

func RespondJSON(
	w http.ResponseWriter,
	status int,
	response models.HealthResponse,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}
