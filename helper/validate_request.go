package helper

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError map[string]string

var validate *validator.Validate

func ValidateRequest(value interface{}) ValidationError {
	validate = validator.New()

	// register tag name to be validated instead using field name
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]

		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}

		return name
	})

	err := validate.Struct(value)

	if err != nil {
		return validatorErrorMessage(err)
	}

	return nil
}

func validatorErrorMessage(validationError error) ValidationError {
	// define variable to store error messages and error fields
	errFields := make(map[string]string)

	// make error message for each invalid field
	for _, err := range validationError.(validator.ValidationErrors) {
		// append error message to the map, where the key is field name and value is an error desctiption
		errFields[err.Field()] = msgForTag(err)
	}

	return errFields
}

func msgForTag(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}

	return fieldError.Error()
}
