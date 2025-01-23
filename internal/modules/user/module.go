package user

import (
	"go-codebase/internal/modules/user/delivery/httphandler"
	"go-codebase/pkg/shared/domain"
)

type module struct {
	handlers map[string]httphandler.HTTPHandler
}

func NewUserModule() *module {
	return &module{
		handlers: map[string]httphandler.HTTPHandler{
			domain.V1: httphandler.NewHTTPHandler(),
		},
	}
}

func (m *module) GetHTTPHandler(version string) httphandler.HTTPHandler {
	if handler, exists := m.handlers[version]; exists {
		return handler
	}
	return nil
}

func (m *module) GetPath() string {
	return "users"
}
