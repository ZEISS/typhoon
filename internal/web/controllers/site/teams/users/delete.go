package teams

import (
	"context"

	"github.com/zeiss/typhoon/internal/web/ports"

	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
)

// TeamUserDeleteControllerImpl ...
type TeamUserDeleteControllerImpl struct {
	team  adapters.GothTeam
	store ports.Datastore
	htmx.DefaultController
}

// NewTeamDeleteController ...
func NewTeamDeleteController(store ports.Datastore) *TeamUserDeleteControllerImpl {
	return &TeamUserDeleteControllerImpl{
		team:  adapters.GothTeam{},
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
	return p.Redirect("/site/teams")
}
