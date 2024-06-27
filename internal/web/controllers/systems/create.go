package systems

import (
	"context"

	"github.com/go-playground/validator/v10"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// validat
var validate *validator.Validate

// CreateSystemControllerImpl ...
type CreateSystemControllerImpl struct {
	System models.System
	store  ports.Datastore
	htmx.DefaultController
}

// NewCreateSystemController ...
func NewCreateSystemController(store ports.Datastore) *CreateSystemControllerImpl {
	return &CreateSystemControllerImpl{
		System:            models.System{},
		DefaultController: htmx.DefaultController{},
		store:             store,
	}
}

// Error ...
func (l *CreateSystemControllerImpl) Error(err error) error {
	return toasts.RenderToasts(
		l.Ctx(),
		toasts.Toasts(
			toasts.ToastsProps{},
			toasts.ToastAlertError(
				toasts.ToastProps{},
				htmx.Text(err.Error()),
			),
		),
	)
}

// Prepare ...
func (l *CreateSystemControllerImpl) Prepare() error {
	validate = validator.New()

	err := l.BindBody(&l.System)
	if err != nil {
		return err
	}

	err = validate.StructCtx(l.Context(), &l.System)
	if err != nil {
		return err
	}

	return nil
}

// Post ...
func (l *CreateSystemControllerImpl) Post() error {
	err := l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateSystem(ctx, &l.System)
	})
	if err != nil {
		return err
	}

	return l.Redirect("/systems")
}
