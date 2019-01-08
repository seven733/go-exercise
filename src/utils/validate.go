package utils

import (
	validator "gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func ValidateBody(data interface{}) (string, error) {
	validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		if len(validationErrors) > 0 {
			errDetail := validationErrors[0]
			return errDetail.Namespace(), err
		}
		return "", err
	}
	return "", err
}
