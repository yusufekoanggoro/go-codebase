package base

import (
	"go-codebase/pkg/database/sql"
	"go-codebase/pkg/logger"
)

type ModuleParam struct {
	Postgres *sql.SQLDatabase
	Logger   *logger.Logger
}
