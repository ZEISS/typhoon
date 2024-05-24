package skgs

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// CreateSkgsControllerImpl ...
type CreateSkgsControllerImpl struct {
	ports.Operators
	htmx.DefaultController
}

// NewCreateSkgsController ...
func NewCreateSkgsController(db ports.Operators) *CreateSkgsControllerImpl {
	return &CreateSkgsControllerImpl{db, htmx.DefaultController{}}
}

// Post ...
func (l *NewSkgsControllerImpl) Post() error {
	return nil
}
