package factory

import (
	"go-codebase/internal/factory/base"
	"go-codebase/internal/modules/user"
)

func NewModuleFactory(param *base.ModuleParam) []ModuleFactory {
	modules := []ModuleFactory{
		user.NewUserModule(param),
	}

	return modules
}
