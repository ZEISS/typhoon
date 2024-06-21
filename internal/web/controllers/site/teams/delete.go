package teams

import (
	"context"

	"github.com/zeiss/typhoon/internal/web/ports"

	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
)

// TeamDeleteControllerImpl ...
type TeamDeleteControllerImpl struct {
	team  adapters.GothTeam
	store ports.Datastore
	htmx.DefaultController
}

// NewTeamDeleteController ...
func NewTeamDeleteController(store ports.Datastore) *TeamDeleteControllerImpl {
	return &TeamDeleteControllerImpl{
		team:  adapters.GothTeam{},
		store: store,
	}
}

// Prepare ...
func (p *TeamDeleteControllerImpl) Prepare() error {
	err := p.BindParams(&p.team)
	if err != nil {
		return err
	}

	return p.store.ReadWriteTx(p.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteTeam(ctx, &p.team)
	})
}

// Delete ...
func (p *TeamDeleteControllerImpl) Delete() error {
	return p.Redirect("/site/teams")
}
