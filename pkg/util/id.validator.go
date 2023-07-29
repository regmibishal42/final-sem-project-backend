package util

import (
	"github.com/google/uuid"
)

func IsValidID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

func IsValidIDs(ids []string) bool {
	isValidID := true
	if len(ids) < 1 {
		return false
	}
	for _, id := range ids {
		_, err := uuid.Parse(id)
		if err != nil {
			isValidID = false
		}
	}
	return isValidID
}
