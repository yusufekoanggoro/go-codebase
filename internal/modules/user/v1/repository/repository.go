package repository

import "go-codebase/internal/modules/user/v1/repository/sql"

type Repository interface {
	GetUserRepoSQL() sql.UserRepoSQL
}

type RepositoryImpl struct {
	userRepoSQL sql.UserRepoSQL
}

func NewRepository() Repository {
	return &RepositoryImpl{
		userRepoSQL: sql.NewUserRepoSQL(),
	}
}

func (r *RepositoryImpl) GetUserRepoSQL() sql.UserRepoSQL {
	return r.userRepoSQL
}
