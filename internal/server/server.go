package server

import (
	"fmt"
	handler "forum/internal/handlers"
	"forum/internal/repository"
	"forum/internal/service"
	"forum/internal/storage"
	"forum/pkg/config"
	"net/http"
)

type App struct {
	cfg config.Config
}

func NewApp(cfg config.Config) *App {
	return &App{cfg: cfg}
}

func (app *App) Run() error {
	logger := handler.NewLogger()
	db, err := storage.NewSqlite(app.cfg)
	if err != nil {
		return err
	}
	logger.Info("Database successfully connected!")

	repo := repository.NewRepo(db)

	logger.Info("Repository working...")

	service := service.NewService(repo)

	logger.Info("Service working...")

	handler := handler.NewHandler(service)

	logger.Info("Handler working...")

	logger.Info("Server successfully started!")
	fmt.Printf("Server running on http://localhost%v\n", app.cfg.Port)

	return http.ListenAndServe(app.cfg.Port, handler.Router())
}
