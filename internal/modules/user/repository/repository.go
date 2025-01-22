package repository

import "go-codebase/internal/modules/user/repository/sql"

type Repository interface {
	GetUserRepoSQL() sql.UserRepoSQL
}

type repositoryImpl struct {
	userRepoSQL sql.UserRepoSQL
}

func NewRepository() Repository {
	return &repositoryImpl{
		userRepoSQL: sql.NewUserRepoSQL(),
	}
}

func (r *repositoryImpl) GetUserRepoSQL() sql.UserRepoSQL {
	return r.userRepoSQL
}
