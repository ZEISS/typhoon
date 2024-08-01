package teams

import (
	"context"

	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/ports"

	"github.com/go-playground/validator/v10"
	htmx "github.com/zeiss/fiber-htmx"
)

const (
	listTeamURL = "/site/teams"
)

var validate *validator.Validate

// CreateTeamControllerImpl ...
type CreateTeamControllerImpl struct {
	team  models.Team
	store ports.Datastore
	htmx.DefaultController
}

// NewCreateTeamController ...
func NewCreateTeamController(store ports.Datastore) *CreateTeamControllerImpl {
	return &CreateTeamControllerImpl{
		store: store,
	}
}

// Prepare ...
func (l *CreateTeamControllerImpl) Prepare() error {
	validate = validator.New()

	err := l.BindBody(&l.team)
	if err != nil {
		return err
	}

	err = validate.Struct(&l.team)
	if err != nil {
		return err
	}

	return l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateTeam(ctx, &l.team)
	})
}

// Post ...
func (l *CreateTeamControllerImpl) Post() error {
	return l.Redirect(listTeamURL)
}
