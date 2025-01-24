package app

import (
	"context"
	"fmt"
	"go-codebase/config"
	"go-codebase/internal/factory"
	"go-codebase/internal/factory/base"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	httpServer *fiber.App
	modules    []factory.Module
}

func NewApp(cfg *config.Config) *App {
	param := &base.ModuleParam{
		Postgres: cfg.GetPostgres(),
	}
	modules := factory.NewModuleFactory(param).GetModules()

	httpServer := fiber.New()

	return &App{
		httpServer: httpServer,
		modules:    modules,
	}
}

func (a *App) Shutdown(ctx context.Context) {
	if err := a.httpServer.Shutdown(); err != nil {
		panic(err)
	}
	fmt.Println("Application shutdown completed")
}
