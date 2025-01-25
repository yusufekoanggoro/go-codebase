package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

func (v *Validator) ValidateRequest(data interface{}) error {
	err := v.validate.Struct(data)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}
