package validator

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

// ValidateStruct validates a struct according to its tags
func ValidateStruct(data interface{}, checkRequired bool) []*ErrorResponse {
	var errors []*ErrorResponse
	err := Validator.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if !checkRequired && err.Tag() == "required" {
				continue
			}

			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

// NewValidator creates a new validator
func NewValidator() *validator.Validate {
	newValidator := validator.New()
	return newValidator
}

// Validator is the validator, it needs to be set before use
var Validator *validator.Validate
