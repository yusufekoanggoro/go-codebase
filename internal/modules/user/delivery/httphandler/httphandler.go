package httphandler

import "github.com/gofiber/fiber/v2"

type HTTPHandler interface {
	RegisterRoutes(fiber.Router)
}

type httpHandlerImpl struct {
}

func NewHTTPHandler() HTTPHandler {
	return &httpHandlerImpl{}
}

func (h *httpHandlerImpl) RegisterRoutes(route fiber.Router) {
	user := route.Group("")            // User route group
	user.Post("/create", h.createUser) // Create user route
}

func (h *httpHandlerImpl) createUser(c *fiber.Ctx) error {
	return nil
}
