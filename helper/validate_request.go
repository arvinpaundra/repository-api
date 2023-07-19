package helper

import (
	"fmt"
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
		// append error message to the map, where the key is field name and value is an error description
		errFields[err.Field()] = msgForTag(err)
	}

	return errFields
}

func msgForTag(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "Wajib diisi"
	case "email":
		return "Email tidak valid"
	case "hexcolor":
		return "Kode warna tidak valid"
	case "base64rawurl":
		return "Format harus berupa base64 raw url encoding"
	case "min":
		min := fieldError.Param()
		return fmt.Sprintf("Minimal panjang %s karakter", min)
	case "max":
		max := fieldError.Param()
		return fmt.Sprintf("Maksimal panjang %s karakter", max)
	case "eqfield":
		return "Input tidak sama"
	case "numeric":
		return "Hanya boleh berupa angka"
	case "alpha":
		return "Hanya boleh berupa huruf"
	default:
		return fieldError.Error()
	}
}
