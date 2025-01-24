package interfaces

import "github.com/gofiber/fiber/v2"

type FiberRestDelivery interface {
	RegisterRoutes(route fiber.Router)
}
