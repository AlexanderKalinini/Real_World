package app

import (
	"github.com/joho/godotenv"
	"net/http"
	"rwa/internal/provider"
)

func GetApp() http.Handler {
	err := godotenv.Load(".env")
	if err != nil {
		return nil
	}

	serviceProvider := provider.NewServiceProvider()
	serviceProvider.InitDeps()

	return serviceProvider.GetApiRouter().Router
}
