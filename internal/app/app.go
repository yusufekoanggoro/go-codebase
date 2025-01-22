package app

import (
	"context"
	"fmt"
	"go-codebase/internal/modules"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	httpServer *fiber.App
	modules    []modules.Module
}

func NewApp() *App {
	httpServer := fiber.New()
	modules := modules.NewModules(&modules.ModuleParam{})

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
