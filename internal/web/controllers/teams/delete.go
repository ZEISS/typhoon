package teams

import (
	"context"

	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
)

// TeamDeleteControllerImpl ...
type TeamDeleteControllerImpl struct {
	store ports.Datastore
	htmx.DefaultController
	team models.Team
}

// NewTeamDeleteController ...
func NewTeamDeleteController(store ports.Datastore) *TeamDeleteControllerImpl {
	return &TeamDeleteControllerImpl{store: store}
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
	return p.Redirect("/teams")
}
