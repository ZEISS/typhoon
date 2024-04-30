package models

import "strings"

// AccountPublicKey ...
type AccountPublicKey string

// String ...
func (a AccountPublicKey) String() string {
	return strings.TrimSpace(string(a))
}
