package strutil

import "github.com/google/uuid"

// GetUUID get uuid string
func GetUUID() string {
	return uuid.New().String()
}
