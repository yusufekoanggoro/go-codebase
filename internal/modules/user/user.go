package user

import (
	"go-codebase/internal/modules/user/delivery/httphandler"
	"go-codebase/pkg/shared/domain"
)

type Module interface {
	HTTPHandler(version string) httphandler.HTTPHandler
	Path() string
}

type moduleImpl struct {
	handlers map[string]httphandler.HTTPHandler
}

func NewUserModule() Module {
	return &moduleImpl{
		handlers: map[string]httphandler.HTTPHandler{
			domain.V1: httphandler.NewHTTPHandler(),
		},
	}
}

func (m *moduleImpl) HTTPHandler(version string) httphandler.HTTPHandler {
	if handler, exists := m.handlers[version]; exists {
		return handler
	}
	return nil
}

func (m *moduleImpl) Path() string {
	return "users"
}
