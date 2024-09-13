package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// FormatValidationError formats the binding errors to be more user-friendly.
func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)

	validationErrors, ok := err.(validator.ValidationErrors)

	fmt.Println("Validation Errors:", validationErrors)
	fmt.Println("Validation ok:", ok)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {

		fmt.Println("Validation Errors:", validationErrors)

		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			tag := fieldError.Tag()
			fmt.Println("Field Error:", fieldName, "Tag:", tag)

			switch tag {
			case "required":
				errors[fieldName] = fieldName + " is required"
			case "email":
				errors[fieldName] = fieldName + " must be a valid email address"
			default:
				errors[fieldName] = "Validation failed on " + tag + " for " + fieldName
			}
		}
	}

	return errors
}

// RespondError creates a formatted error response.
func RespondError(c *gin.Context, statusCode int, message string, err error) gin.H {
	return gin.H{
		"error":   http.StatusText(statusCode),
		"details": FormatValidationError(err),
	}
}
