package teams

import (
	"context"

	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
)

// TeamUserDeleteControllerImpl ...
type TeamUserDeleteControllerImpl struct {
	team  models.Team
	store ports.Datastore
	htmx.DefaultController
}

// NewTeamDeleteController ...
func NewTeamDeleteController(store ports.Datastore) *TeamUserDeleteControllerImpl {
	return &TeamUserDeleteControllerImpl{
		team:  models.Team{},
		store: store,
	}
}

// Prepare ...
func (p *TeamUserDeleteControllerImpl) Prepare() error {
	err := p.BindParams(&p.team)
	if err != nil {
		return err
	}

	return p.store.ReadWriteTx(p.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteTeam(ctx, &p.team)
	})
}

// Delete ...
func (p *TeamUserDeleteControllerImpl) Delete() error {
	return p.Redirect("/teams")
}
