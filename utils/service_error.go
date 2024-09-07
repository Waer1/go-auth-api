package utils

import (
	"encoding/json"
)

// ServiceErr represents an error with an associated HTTP status code and a message map.
type ServiceErr struct {
	StatusCode int               // HTTP status code
	Message    map[string]string // Message contains structured error data
}

// NewServiceErr creates a new ServiceErr with the given status code and message.
func NewServiceErr(statusCode int, message map[string]string) *ServiceErr {
	return &ServiceErr{
		StatusCode: statusCode,
		Message:    message,
	}
}

// Error implements the error interface. It returns the JSON representation of the error message.
func (e *ServiceErr) Error() string {
	msg, err := json.Marshal(e.Message)
	if err != nil {
		return "error marshalling message"
	}
	return string(msg)
}
