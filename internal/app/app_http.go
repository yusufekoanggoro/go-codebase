package app

import (
	"fmt"
	"go-codebase/config"
	"go-codebase/pkg/shared/domain"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (a *App) ServeHTTP() {
	httpServer := a.httpServer.Group("/code-base")
	httpServer.Get("", func(c *fiber.Ctx) error {
		if err := c.JSON(fiber.Map{"status": "Service up and running"}); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to respond")
		}
		return nil
	})

	for _, m := range a.modules {
		if h := m.GetHTTPHandler(domain.V1); h != nil {
			v1Path := fmt.Sprintf("/%s/%s", m.GetPath(), domain.V1)
			v1Group := httpServer.Group(v1Path)
			h.RegisterRoutes(v1Group)
		}
	}

	address := fmt.Sprintf(":%s", config.GlobalEnv.HTTPPort)

	if err := a.httpServer.Listen(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
