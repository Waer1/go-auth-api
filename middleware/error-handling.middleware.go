package middleware

import (
	"api-auth/utils" // Assuming utils is the package where ServiceErr is defined
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Execute the handler

		// After the handler execution, check for errors
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				// Check if the error is a ServiceErr and handle accordingly
				if serr, ok := e.Err.(*utils.ServiceErr); ok {
					c.JSON(serr.StatusCode, gin.H{
						"error":   http.StatusText(serr.StatusCode),
						"details": serr.Message,
					})
					c.Abort() // Prevent processing of further middleware/handlers
					return
				}
			}

			// If the error is not a ServiceErr, handle it as a generic internal error
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "An unexpected error occurred",
			})
			c.Abort()
		}
	}
}
