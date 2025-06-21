package app

import (
	userContr "rwa/internal/controller/user"
	"rwa/internal/database"
	userRepo "rwa/internal/infractructure/repository/user"
	"rwa/internal/route"
	"rwa/internal/usecase/user"
)

type App struct {
	Router *route.Router
	DB     *database.Sql
}

func NewApp() (*App, error) {

	db, err := database.InitSqlDB()
	if err != nil {
		return nil, err
	}

	repo := userRepo.NewRepository(db.DB)
	useCase := user.NewUseCase(repo)
	userController := userContr.NewUserController(useCase)
	apiRouter := route.NewApiRouter(userController)

	app := &App{
		Router: apiRouter,
		DB:     db,
	}

	return app, nil
}

func (a *App) Close() error {
	return a.DB.DB.Close()
}
