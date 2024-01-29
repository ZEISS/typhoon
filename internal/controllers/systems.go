package controllers

import "github.com/zeiss/typhoon/internal/ports"

// Systems is the controller that uses the Systems interface.
type Systems struct{}

// New returns a new instance of Systems.
func NewSystemsContoller(repo ports.Repositories) *Systems {
	return &Systems{}
}
