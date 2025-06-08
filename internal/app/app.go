package app

import (
	"net/http"
	"rwa/internal/provider"
)

func GetApp() http.Handler {

	serviceProvider := provider.NewServiceProvider()
	serviceProvider.InitDeps()

	return serviceProvider.GetApiRouter().Router
}
