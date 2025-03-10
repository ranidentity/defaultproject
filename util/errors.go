package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func HandleValidationErrors(err error, trans map[string]string) map[string]string {
	validationErrors := err.(validator.ValidationErrors)
	customErrors := make(map[string]string)

	// Iterate over each validation error and map it to the custom error message
	for _, fieldError := range validationErrors {
		fieldName := fieldError.Field()
		switch fieldError.Tag() {
		case "required":
			// Translate error message for required fields
			if message, exists := trans[fieldName+".required"]; exists {
				customErrors[fieldName] = message
			} else {
				customErrors[fieldName] = fmt.Sprintf("%s is required", fieldName)
			}
		default:
			customErrors[fieldName] = fmt.Sprintf("Invalid value for %s", fieldName)
		}
	}
	return customErrors
}
