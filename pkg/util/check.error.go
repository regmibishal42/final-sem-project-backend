package util

import "strings"

func IsDuplicateKeyError(err error) bool {
	// Customize this logic based on the actual error message you receive
	// Look for specific strings or patterns that indicate a duplicate key error
	return strings.Contains(err.Error(), "duplicate key") ||
		strings.Contains(err.Error(), "unique constraint")
}
