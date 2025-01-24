package factory

import (
	"go-codebase/internal/factory/interfaces"
)

type Module interface {
	GetHTTPHandler(version string) interfaces.FiberRestDelivery
	GetPath() string
}

type ModuleFactory struct {
	modules []Module
}
