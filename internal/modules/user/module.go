package user

import (
	"go-codebase/internal/factory/base"
	"go-codebase/internal/factory/interfaces"
	deliveryV1 "go-codebase/internal/modules/user/v1/delivery/httphandler"
	"go-codebase/pkg/shared/domain"
)

type Module struct {
	v1 struct {
		httpHandler *deliveryV1.HTTPHandler
	}
}

func NewUserModule(param *base.ModuleParam) *Module {
	var module Module
	module.v1.httpHandler = deliveryV1.NewHTTPHandler()
	return &module
}

func (m *Module) GetHTTPHandler(version string) interfaces.FiberRestDelivery {
	switch version {
	case domain.V1:
		return m.v1.httpHandler
	}
	return nil
}

func (m *Module) GetPath() string {
	return "users"
}
