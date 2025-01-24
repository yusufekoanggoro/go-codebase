package factory

import (
	"go-codebase/internal/factory/base"
	"go-codebase/internal/factory/interfaces"
)

type ModuleFactory interface {
	GetHTTPHandler(version string) interfaces.FiberRestDelivery
	GetPath() string
}

type ModuleFactoryImpl struct {
	param *base.ModuleParam
}
