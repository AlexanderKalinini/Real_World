package route

import (
	"github.com/gorilla/mux"
	"rwa/internal/controller/user"
)

type Router struct {
	Router         *mux.Router
	userController *user.CreateUserController
}

func NewApiRouter(userController *user.CreateUserController) *Router {
	router := &Router{
		Router:         mux.NewRouter(),
		userController: userController,
	}
	router.registerRoutes()
	return router
}

func (r *Router) registerRoutes() {
	r.Router.HandleFunc("/users", r.userController.Create).Methods("POST")
}
