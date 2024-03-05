package adapter

import (
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

var _ openapi.StrictServerInterface = (*Handlers)(nil)

// Handlers ...
type Handlers struct {
	openapi.Unimplemented
}

// NewHandlers ...
func NewHandlers() *Handlers {
	return &Handlers{}
}
