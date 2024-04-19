package adapters

import (
	"github.com/zeiss/typhoon/internal/api/ports"
	"github.com/zeiss/typhoon/internal/build"
)

var _ ports.Build = (*Build)(nil)

// Build ...
type Build struct{}

// NewBuild ...
func NewBuild() *Build {
	return &Build{}
}

// Version ...
func (b *Build) Version() (string, error) {
	return build.Version(), nil
}

// Build ...
func (b *Build) Build() (string, error) {
	return build.Build(), nil
}

// Date ...
func (b *Build) Date() (string, error) {
	return build.Date(), nil
}
