package errors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/realwebdev/blog/internal/providers/validation"
)

func FromValidation(err error) map[string]string {
	errorsList := make(map[string]string)
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			errorsList[strings.ToLower(fieldError.Field())] = GetErrorMsg(fieldError.Tag())
		}
	}
	return errorsList
}

func GetErrorMsg(tag string) string {
	if msg, ok := validation.ErrorMessages()[tag]; ok {
		return msg
	}
	return "Invalid value"
}
