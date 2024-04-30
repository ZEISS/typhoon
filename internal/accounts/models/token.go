package models

import "strings"

// AccountToken ...
type AccountToken string

// String ...
func (a AccountToken) String() string {
	return strings.TrimSpace(string(a))
}
