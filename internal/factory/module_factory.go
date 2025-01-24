package factory

import (
	"go-codebase/internal/factory/base"
	"go-codebase/internal/modules/user"
)

func NewModuleFactory(param *base.ModuleParam) *ModuleFactory {
	modules := []Module{
		user.NewUserModule(param),
	}

	return &ModuleFactory{
		modules: modules,
	}
}

func (mf *ModuleFactory) GetModules() []Module {
	return mf.modules
}
