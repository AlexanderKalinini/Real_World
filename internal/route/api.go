package route

import (
	"github.com/gorilla/mux"
	"rwa/internal/controller/user"
	"rwa/internal/middleware"
)

type Router struct {
	Router         *mux.Router
	userController *user.UsersController
}

func NewApiRouter(userController *user.UsersController) *Router {
	router := &Router{
		Router:         mux.NewRouter(),
		userController: userController,
	}
	router.registerRoutes()
	return router
}

func (r *Router) registerRoutes() {
	r.Router.HandleFunc("/users", middleware.ErrorHandler(r.userController.Create)).Methods("POST")
}
