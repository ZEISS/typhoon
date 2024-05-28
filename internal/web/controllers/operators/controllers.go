package operators

import (
	"github.com/zeiss/typhoon/internal/web/controllers/operators/sysaccount"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// NewUpdateSystemAccountController ...
func NewUpdateSystemAccountController(db ports.Repository) *sysaccount.UpdateSystemAccountControllerImpl {
	return sysaccount.NewCreateSkgsController(db)
}
