package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BindJSONAndValidate attempts to bind a JSON payload to a model and validates it.
// It returns true if binding and validation are successful, or false and handles the error response.
func BindJSONAndValidate(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		fmt.Println("Error binding JSON:", err)
		errDetails := FormatValidationError(err)
		c.Error(
			NewServiceErr(http.StatusBadRequest, errDetails),
		) // Add error to Gin Context
		c.Status(http.StatusBadRequest) // Set status code to 400 Bad Request
		return false
	}
	return true
}
