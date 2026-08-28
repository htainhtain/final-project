package sqlhealth

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"server/models"
	"server/respond"
	"time"

	_ "github.com/microsoft/go-mssqldb/azuread"
)

// --------------------------------------------------
// Azure SQL
// --------------------------------------------------

func SqlHealth(w http.ResponseWriter, r *http.Request) {
	server := os.Getenv("SQL_SERVER")
	database := os.Getenv("SQL_SERVER_DATABASE")

	connectionString := fmt.Sprintf(
		"sqlserver://%s?database=%s&fedauth=ActiveDirectoryManagedIdentity&encrypt=true&TrustServerCertificate=false",
		server,
		database,
	)

	if connectionString == "" {
		respond.RespondJSON(w, http.StatusServiceUnavailable, models.HealthResponse{
			Status:  "error",
			Service: "azure-sql",
			Message: "SQL_CONNECTION_STRING is not configured",
		})
		return
	}

	db, err := sql.Open("azuresql", connectionString)
	fmt.Println("err: ", err)
	if err != nil {
		respond.RespondJSON(w, http.StatusServiceUnavailable, models.HealthResponse{
			Status:  "error",
			Service: "azure-sql",
			Message: "Could not open SQL connection",
		})
		return
	}

	defer db.Close()

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		fmt.Println("err: ", err)
		respond.RespondJSON(w, http.StatusServiceUnavailable, models.HealthResponse{
			Status:  "error",
			Service: "azure-sql",
			Message: "Azure SQL is not reachable",
		})
		return
	}

	respond.RespondJSON(w, http.StatusOK, models.HealthResponse{
		Status:  "ok",
		Service: "azure-sql",
		Message: "Azure SQL is working",
	})
}
