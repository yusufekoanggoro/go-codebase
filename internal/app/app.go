package app

import (
	"context"
	"go-codebase/config"
	"go-codebase/internal/factory"
	"go-codebase/internal/factory/base"
	"go-codebase/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	httpServer *fiber.App
	modules    []factory.Module
	logger     logger.Logger
}

func NewApp(cfg *config.Config) *App {
	param := &base.ModuleParam{
		Postgres: cfg.GetPostgres(),
		Logger:   cfg.GetLogger(),
	}
	modules := factory.NewModuleFactory(param).GetModules()

	httpServer := fiber.New()

	return &App{
		httpServer: httpServer,
		modules:    modules,
		logger:     cfg.GetLogger(),
	}
}

func (a *App) Shutdown(ctx context.Context) {
	if err := a.httpServer.Shutdown(); err != nil {
		panic(err)
	}
	a.logger.Info("Application shutdown completed", "App.Shutdown()", "appshutdown")
}
