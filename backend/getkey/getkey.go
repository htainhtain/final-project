package getkey

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

func GetKey(vaultName, secretName string) (string, error) {
	ctx := context.Background()

	if vaultName == "" {
		return "", fmt.Errorf("KEY_VAULT_NAME is required")
	}

	if secretName == "" {
		return "", fmt.Errorf("SECRET_NAME is required")
	}

	vaultURL := fmt.Sprintf(
		"https://%s.vault.azure.net/",
		vaultName,
	)

	if vaultURL == "" {
		return "", fmt.Errorf("Invalid key vault URL")
	}

	fmt.Println("beofre managed identity")
	// Use the managed identity assigned to the ACI.
	credential, err := azidentity.NewManagedIdentityCredential(nil)
	if err != nil {
		return "", fmt.Errorf("failed to create managed identity credential: %v", err)
	}
	fmt.Println("after managed identity")

	client, err := azsecrets.NewClient(
		vaultURL,
		credential,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create managed identity credential: %v", err)
	}

	secret, err := client.GetSecret(
		ctx,
		secretName,
		"",
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to get secret: %v", err)
	}

	if secret.Value == nil {
		return "", fmt.Errorf("secret value is nil")
	}

	return *secret.Value, nil
}
