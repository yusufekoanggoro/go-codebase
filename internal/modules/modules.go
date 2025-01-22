package modules

import (
	"go-codebase/internal/modules/user"
	"go-codebase/internal/modules/user/delivery/httphandler"
)

type Module interface {
	HTTPHandler(version string) httphandler.HTTPHandler
	Path() string
}

type ModuleParam struct {
}

func NewModules(param *ModuleParam) []Module {
	return []Module{
		user.NewUserModule(),
	}
}
