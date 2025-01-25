package httphandler

import "github.com/gofiber/fiber/v2"

type HTTPHandler struct {
}

func NewHTTPHandler() *HTTPHandler {
	return &HTTPHandler{}
}

func (h *HTTPHandler) RegisterRoutes(route fiber.Router) {
	user := route.Group("")            // User route group
	user.Post("/create", h.createUser) // Create user route
}

func (h *HTTPHandler) createUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Status": "Created",
	})
}
