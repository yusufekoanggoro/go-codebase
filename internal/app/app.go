package app

import (
	"context"
	"fmt"
	"go-codebase/internal/factory"
	"go-codebase/internal/factory/base"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	httpServer *fiber.App
	modules    []factory.ModuleFactory
}

func NewApp() *App {
	httpServer := fiber.New()
	modules := factory.NewModuleFactory(&base.ModuleParam{})

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
