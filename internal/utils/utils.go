package utils

import "github.com/google/uuid"

// PtrInt returns a pointer to an int.
func PtrInt(i int) *int {
	return &i
}

// PtrUUID returns a pointer to a UUID.
func PtrUUID(u uuid.UUID) *uuid.UUID {
	return &u
}

// StrPtr returns a pointer to a string.
func StrPtr(s string) *string {
	return &s
}
