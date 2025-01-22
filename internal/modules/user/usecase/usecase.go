package usecase

import (
	"context"
)

// UserUseCase defines the contract for user-related use cases.
type UserUseCase interface {
	GetUsers(ctx context.Context) error
}

// userUseCaseImpl is the implementation of UserUseCase.
type userUseCaseImpl struct{}

// NewUserUsecase creates a new instance of UserUseCase.
func NewUserUseCase() UserUseCase {
	return &userUseCaseImpl{}
}

// GetUsers handles the business logic to retrieve user data.
func (uc *userUseCaseImpl) GetUsers(ctx context.Context) error {
	return nil
}
