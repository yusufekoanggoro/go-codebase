package user

import (
	"go-codebase/internal/factory/base"
	"go-codebase/internal/factory/interfaces"
	deliveryV1 "go-codebase/internal/modules/user/v1/delivery/httphandler"
	"go-codebase/pkg/shared/domain"
)

type module struct {
	v1 struct {
		httpHandler *deliveryV1.HTTPHandler
	}
}

func NewUserModule(param *base.ModuleParam) *module {
	var module module
	module.v1.httpHandler = deliveryV1.NewHTTPHandler()
	return &module
}

func (m *module) GetHTTPHandler(version string) interfaces.FiberRestDelivery {
	switch version {
	case domain.V1:
		return m.v1.httpHandler
	}
	return nil
}

func (m *module) GetPath() string {
	return "users"
}
