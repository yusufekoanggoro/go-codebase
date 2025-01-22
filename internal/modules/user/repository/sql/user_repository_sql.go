package sql

import (
	"context"
)

type UserRepoSQL interface {
	Create(ctx context.Context) error
}

type userRepoSQLImpl struct {
}

func NewUserRepoSQL() UserRepoSQL {
	return &userRepoSQLImpl{}
}

func (r *userRepoSQLImpl) Create(ctx context.Context) error {
	return nil
}
