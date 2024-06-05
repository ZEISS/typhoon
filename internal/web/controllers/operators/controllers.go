package operators

import (
	"github.com/zeiss/typhoon/internal/web/controllers/operators/sysaccount"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// NewUpdateSystemAccountController ...
func NewUpdateSystemAccountController(store ports.Datastore) *sysaccount.UpdateSystemAccountControllerImpl {
	return sysaccount.NewCreateSkgsController(store)
}
