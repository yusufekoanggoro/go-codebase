package usecase

import (
	"context"
)

// UserUseCase defines the contract for user-related use cases.
type UserUseCase interface {
	GetUsers(ctx context.Context) error
}

// UserUseCaseImpl is the implementation of UserUseCase.
type UserUseCaseImpl struct{}

// NewUserUsecase creates a new instance of UserUseCase.
func NewUserUseCase() UserUseCase {
	return &UserUseCaseImpl{}
}

// GetUsers handles the business logic to retrieve user data.
func (uc *UserUseCaseImpl) GetUsers(ctx context.Context) error {
	return nil
}
