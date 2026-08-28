package keyvault

import (
	"log"
	"net/http"
	"os"
	"server/getkey"
	"server/models"
	"server/respond"
)

// --------------------------------------------------
// Key Vault
// --------------------------------------------------

func KeyVaultHealth(w http.ResponseWriter, r *http.Request) {
	vaultName := os.Getenv("AZURE_KEY_VAULT_NAME")
	secretName := os.Getenv("AZURE_KEY_VAULT_SECRET")

	key, err := getkey.GetKey(vaultName, secretName)
	if err != nil {
		log.Fatalf("failed to get secret: %v", err)
	}

	respond.RespondJSON(w, http.StatusOK, models.HealthResponse{
		Status:  "ok",
		Service: "keyvault",
		Message: "Secret key value is " + key,
	})
}
