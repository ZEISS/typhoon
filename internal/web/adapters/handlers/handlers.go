package handlers

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/web/controllers/accounts"
	pa "github.com/zeiss/typhoon/internal/web/controllers/accounts/partials"
	"github.com/zeiss/typhoon/internal/web/controllers/dashboard"
	"github.com/zeiss/typhoon/internal/web/controllers/login"
	"github.com/zeiss/typhoon/internal/web/controllers/me"
	"github.com/zeiss/typhoon/internal/web/controllers/operators"
	oskgs "github.com/zeiss/typhoon/internal/web/controllers/operators/skgs"
	ot "github.com/zeiss/typhoon/internal/web/controllers/operators/tokens"
	"github.com/zeiss/typhoon/internal/web/controllers/systems"
	"github.com/zeiss/typhoon/internal/web/controllers/users"
	"github.com/zeiss/typhoon/internal/web/controllers/users/credentials"
	pu "github.com/zeiss/typhoon/internal/web/controllers/users/partials"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var _ ports.Handlers = (*handlers)(nil)

type handlers struct {
	store ports.Datastore
}

// NewHandlers ...
func NewHandlers(store ports.Datastore) *handlers {
	return &handlers{store}
}

// Login ...
func (h *handlers) Login() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return login.NewIndexLoginController()
	})
}

// Dashboard ...
func (h *handlers) Dashboard() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return dashboard.NewIndexDashboardController()
	})
}

// Me ...
func (h *handlers) Me() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return me.NewMeController(h.store)
	})
}

// ListOperators ...
func (h *handlers) ListOperators() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return operators.NewListOperatorsController(h.store)
	})
}

// NewOperator ...
func (h *handlers) NewOperator() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return operators.NewOperatorController(h.store)
	})
}

// CreateOperator ...
func (h *handlers) CreateOperator() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return operators.NewOperatorController(h.store)
	})
}

// ShowOperator ...
func (h *handlers) ShowOperator() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return operators.NewShowOperatorController(h.store)
	})
}

// TokenOperator ...
func (h *handlers) TokenOperator() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return ot.NewIndexOperatorTokenController(h.store)
	})
}

// DeleteOperator ...
func (h *handlers) DeleteOperator() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return operators.NewDeleteOperatorController(h.store)
	})
}

// ListAccounts ...
func (h *handlers) ListAccounts() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return accounts.NewListAccountsController(h.store)
	})
}

// NewAccount ...
func (h *handlers) NewAccount() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return accounts.NewAccountController(h.store)
	})
}

// ListUsers ...
func (h *handlers) ListUsers() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return users.NewListUsersController(h.store)
	})
}

// CreateAccount ...
func (h *handlers) CreateAccount() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return accounts.NewCreateController(h.store)
	})
}

// ShowAccount ...
func (h *handlers) ShowAccount() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return accounts.NewShowAccountController(h.store)
	})
}

// DeleteAccount ...
func (h *handlers) DeleteAccount() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return accounts.NewDeleteAccountController(h.store)
	})
}

// NewOperatorSkg ...
func (h *handlers) NewOperatorSkg() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return oskgs.NewSkgsController(h.store)
	})
}

// CreateOperatorSkg ...
func (h *handlers) CreateOperatorSkg() fiber.Handler {
	return htmx.NewHxControllerHandler(
		func() htmx.Controller {
			return oskgs.NewCreateSkgsController(h.store)
		})
}

// OperatorSkgsOptions ...
func (h *handlers) OperatorSkgsOptions() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return pa.NewOperatorSkgsOptions(h.store)
	})
}

// AccountSksOptions ...
func (h *handlers) AccountSksOptions() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return pu.NewAccountSkgsOptions(h.store)
	})
}

// NewUser ...
func (h *handlers) NewUser() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return users.NewUserController(h.store)
	})
}

// CreateUser ...
func (h *handlers) CreateUser() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return users.NewCreateUserController(h.store)
	})
}

// ShowUser ...
func (h *handlers) ShowUser() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return users.NewShowUserController(h.store)
	})
}

// UserCredentials ...
func (h *handlers) UserCredentials() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return credentials.NewIndexUserCredentialsController(h.store)
	})
}

// DeleteUser ...
func (h *handlers) DeleteUser() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return users.NewDeleteUserController(h.store)
	})
}

// UpdateSystemAccount ...
func (h *handlers) UpdateSystemAccount() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return operators.NewUpdateSystemAccountController(h.store)
	})
}

// GetAccountToken ...
func (h *handlers) GetAccountToken() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return accounts.NewGetAccountTokenController(h.store)
	})
}

// ListSystems ...
func (h *handlers) ListSystems() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return systems.NewListSystemsController(h.store)
	})
}
