package backendHealth

import (
	"net/http"
	"server/models"
	"server/respond"
)

func BackendHealth(w http.ResponseWriter, r *http.Request) {
	respond.RespondJSON(w, http.StatusOK, models.HealthResponse{
		Status:  "ok",
		Service: "backend",
		Message: "pong",
	})
}
