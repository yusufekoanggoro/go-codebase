package modules

import (
	"go-codebase/internal/modules/user"
	"go-codebase/internal/modules/user/delivery/httphandler"
)

type Module interface {
	GetPath() string
	GetHTTPHandler(version string) httphandler.HTTPHandler
}

type ModuleParam struct {
}

func NewRegistryModule(param *ModuleParam) []Module {
	modules := []Module{
		user.NewUserModule(),
	}

	return modules
}
