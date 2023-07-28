package model

import (
	"net/http"
)

func CreateValidationError(message string, field string) *ValidationError {
	var msgs []string
	msgs = append(msgs, message)
	return &ValidationError{
		Message: message,
		Code:    http.StatusNotAcceptable,
		// Field:   []scalar.InvalidData{map[string][]string{field: msgs}},
	}
}
