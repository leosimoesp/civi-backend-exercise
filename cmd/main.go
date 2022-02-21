package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/leosimoesp/civi-backend-exercise/api"
	"github.com/leosimoesp/civi-backend-exercise/config"
	"github.com/leosimoesp/civi-backend-exercise/internal/app/repo"
	"github.com/leosimoesp/civi-backend-exercise/internal/app/service"
)

func init() {
	checkEnvVars([]string{config.Port, config.FilePath})
	repo.LoadPoints()
}

func main() {
	port := os.Getenv(config.Port)
	address := fmt.Sprint(":", port)

	log.Printf("Running server at port %v\n", port)

	//GET /api/points
	repo := repo.NewRepo()
	apiService := service.NewCartesianService(repo)

	routes := api.NewApiRoutes(apiService)
	routes.HandleGetPointsWithinDistance()

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal(err)
	}
}

// validate all env vars by name into os path
func checkEnvVars(envVarsNames []string) {
	for _, v := range envVarsNames {
		if os.Getenv(v) == "" {
			log.Panicf("Env var %s is mandatory at path app", v)
		}
	}
}
