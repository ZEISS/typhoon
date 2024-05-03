package utils

import (
	"time"

	"github.com/google/uuid"
)

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

// PtrStr returns a pointer to a string.
func PtrStr(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

// PtrTime returns a pointer to a time.
func PtrTime(t time.Time) *time.Time {
	return &t
}

// TimePtr returns a pointer to a time.
func TimePtr(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}

	return *t
}

// PtrUUID returns a pointer to a UUID.
func UUIDPtr(u uuid.UUID) *uuid.UUID {
	return &u
}
