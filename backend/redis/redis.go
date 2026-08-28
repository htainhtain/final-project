package redis

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"server/getkey"
	"server/models"
	"server/respond"
	"time"

	redis "github.com/redis/go-redis/v9"
)

func RedisHealth(w http.ResponseWriter, r *http.Request) {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	vaultName := os.Getenv("AZURE_KEY_VAULT_NAME")
	secretName := os.Getenv("AZURE_KEY_VAULT_REDIS_SECRET")

	redisPassword := os.Getenv("REDIS_PASSWORD")

	redisPassword, err := getkey.GetKey(vaultName, secretName)
	if err != nil {
		fmt.Println(err)
		respond.RespondJSON(w, http.StatusServiceUnavailable, models.HealthResponse{
			Status:  "error",
			Service: "azure-sql",
			Message: "failed to get secret",
		})
		return
	}

	fmt.Println("redis password: ", redisPassword)

	redisURL := fmt.Sprintf(
		"redis://:%s@%s:%s",
		redisPassword,
		redisHost,
		redisPort,
	)

	fmt.Println("redisURL: ", redisURL)

	if redisURL == "" {
		respond.RespondJSON(w, http.StatusServiceUnavailable, models.HealthResponse{
			Status:  "error",
			Service: "redis",
			Message: "REDIS_URL is not configured",
		})
		return
	}

	options, err := redis.ParseURL(redisURL)
	if err != nil {
		respond.RespondJSON(w, http.StatusServiceUnavailable, models.HealthResponse{
			Status:  "error",
			Service: "redis",
			Message: "Invalid Redis URL",
		})
		return
	}

	client := redis.NewClient(options)
	defer client.Close()

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		fmt.Println("err: ", err)
		respond.RespondJSON(w, http.StatusServiceUnavailable, models.HealthResponse{
			Status:  "error",
			Service: "redis",
			Message: "Redis is not reachable",
		})
		return
	}

	respond.RespondJSON(w, http.StatusOK, models.HealthResponse{
		Status:  "ok",
		Service: "redis",
		Message: "Redis is working",
	})
}
