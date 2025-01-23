package app

import (
	"fmt"
	"go-codebase/pkg/shared/domain"

	"github.com/gofiber/fiber/v2"
)

func (a *App) ServeHTTP() {
	httpServer := a.httpServer.Group("/code-base")
	httpServer.Get("", func(c *fiber.Ctx) error {
		return c.JSON(200, "Service up and running")
	})

	for _, m := range a.modules {
		if h := m.GetHTTPHandler(domain.V1); h != nil {
			v1Path := fmt.Sprintf("/%s/%s", m.GetPath(), domain.V1)
			v1Group := httpServer.Group(v1Path)
			h.RegisterRoutes(v1Group)
		}
	}

	if err := a.httpServer.Listen(fmt.Sprintf(":%d", 3000)); err != nil {
		fmt.Println(err)
	}
}
