package provider

import (
	"database/sql"
	"fmt"
	"rwa/internal/controller/user"
	"rwa/internal/database"
	userRepo "rwa/internal/repository/user"
	"rwa/internal/route"
)

type ServiceProvider struct {
	db             *sql.DB
	userRepo       *userRepo.Repository
	userController *user.UsersController
	router         *route.Router
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (s *ServiceProvider) InitDeps() *ServiceProvider {
	s.NewDB()
	s.NewUserRepo()
	s.UserController()
	s.NewApiRouter()
	return s
}

func (s *ServiceProvider) NewDB() *sql.DB {
	d, err := database.InitSqlDB()
	if err != nil {
		return nil
	}
	s.db = d.GetDB()
	return s.db
}

func (s *ServiceProvider) NewUserRepo() *userRepo.Repository {
	s.userRepo = userRepo.NewRepository(s.db)
	return s.userRepo
}

func (s *ServiceProvider) UserController() *user.UsersController {
	s.userController = user.NewUserController(s.userRepo)
	return s.userController
}

func (s *ServiceProvider) NewApiRouter() *route.Router {
	s.router = route.NewApiRouter(s.userController)
	return s.router
}

func (s *ServiceProvider) GetApiRouter() *route.Router {
	return s.router
}

func (s *ServiceProvider) Close() error {
	err := s.db.Close()
	if err != nil {
		return fmt.Errorf("error closing db: %w", err)
	}
	return nil
}
