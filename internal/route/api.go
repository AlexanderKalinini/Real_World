package route

import (
	"github.com/gorilla/mux"
	"rwa/internal/controller"
)

type Router struct {
	Router *mux.Router
}

func NewApiRouter(controllers ...controller.Controller) *Router {
	router := &Router{
		Router: mux.NewRouter(),
	}
	for _, c := range controllers {
		c.Register(router.Router)
	}

	return router
}
