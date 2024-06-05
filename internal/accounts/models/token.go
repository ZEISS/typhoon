package models

import "strings"

// AccountToken ...
type AccountToken []byte

// String ...
func (a AccountToken) String() string {
	return strings.TrimSpace(string(a))
}

// Byte ...
func (a AccountToken) Byte() []byte {
	return a
}
