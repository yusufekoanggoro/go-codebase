package sql

import (
	"context"
)

type UserRepoSQL interface {
	Create(ctx context.Context) error
}

type UserRepoSQLImpl struct {
}

func NewUserRepoSQL() UserRepoSQL {
	return &UserRepoSQLImpl{}
}

func (r *UserRepoSQLImpl) Create(ctx context.Context) error {
	return nil
}
