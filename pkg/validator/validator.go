package validator

import (
	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Validate(s any) error
	// You can add more methods as needed
}

type validatorImpl struct {
	validate *validator.Validate
}

func New() Validator {
	v := validator.New()

	// Example: register custom tag if needed
	// v.RegisterValidation("mytag", yourCustomFunc)

	return &validatorImpl{
		validate: v,
	}
}

func (v *validatorImpl) Validate(s any) error {
	return v.validate.Struct(s)
}
