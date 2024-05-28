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
	"github.com/zeiss/typhoon/internal/web/controllers/users"
	"github.com/zeiss/typhoon/internal/web/controllers/users/credentials"
	pu "github.com/zeiss/typhoon/internal/web/controllers/users/partials"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var _ ports.Handlers = (*handlers)(nil)

type handlers struct {
	db ports.Repository
}

// NewHandlers ...
func NewHandlers(db ports.Repository) *handlers {
	return &handlers{db}
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
		return me.NewMeController()
	})
}

// ListOperators ...
func (h *handlers) ListOperators() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return operators.NewListOperatorsController(h.db)
	})
}

// NewOperator ...
func (h *handlers) NewOperator() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return operators.NewOperatorController(h.db)
	})
}

// CreateOperator ...
func (h *handlers) CreateOperator() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return operators.NewOperatorController(h.db)
	})
}

// ShowOperator ...
func (h *handlers) ShowOperator() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return operators.NewShowOperatorController(h.db)
	})
}

// TokenOperator ...
func (h *handlers) TokenOperator() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return ot.NewIndexOperatorTokenController(h.db)
	})
}

// DeleteOperator ...
func (h *handlers) DeleteOperator() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return operators.NewDeleteOperatorController(h.db)
	})
}

// ListAccounts ...
func (h *handlers) ListAccounts() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return accounts.NewListAccountsController(h.db)
	})
}

// NewAccount ...
func (h *handlers) NewAccount() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return accounts.NewAccountController(h.db)
	})
}

// ListUsers ...
func (h *handlers) ListUsers() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return users.NewListUsersController(h.db)
	})
}

// CreateAccount ...
func (h *handlers) CreateAccount() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return accounts.NewCreateController(h.db)
	})
}

// ShowAccount ...
func (h *handlers) ShowAccount() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return accounts.NewShowAccountController(h.db)
	})
}

// DeleteAccount ...
func (h *handlers) DeleteAccount() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return accounts.NewDeleteAccountController(h.db)
	})
}

// NewOperatorSkg ...
func (h *handlers) NewOperatorSkg() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return oskgs.NewSkgsController(h.db)
	})
}

// CreateOperatorSkg ...
func (h *handlers) CreateOperatorSkg() fiber.Handler {
	return htmx.NewHxControllerHandler(
		func() htmx.Controller {
			return oskgs.NewCreateSkgsController(h.db)
		})
}

// OperatorSkgsOptions ...
func (h *handlers) OperatorSkgsOptions() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return pa.NewOperatorSkgsOptions(h.db)
	})
}

// AccountSksOptions ...
func (h *handlers) AccountSksOptions() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return pu.NewAccountSkgsOptions(h.db)
	})
}

// NewUser ...
func (h *handlers) NewUser() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return users.NewUserController(h.db)
	})
}

// CreateUser ...
func (h *handlers) CreateUser() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return users.NewCreateUserController(h.db)
	})
}

// ShowUser ...
func (h *handlers) ShowUser() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return users.NewShowUserController(h.db)
	})
}

// UserCredentials ...
func (h *handlers) UserCredentials() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return credentials.NewIndexUserCredentialsController(h.db)
	})
}

// DeleteUser ...
func (h *handlers) DeleteUser() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return users.NewDeleteUserController(h.db)
	})
}

// UpdateSystemAccount ...
func (h *handlers) UpdateSystemAccount() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return operators.NewUpdateSystemAccountController(h.db)
	})
}

// GetAccountToken ...
func (h *handlers) GetAccountToken() fiber.Handler {
	return htmx.NewHxControllerHandler(func() htmx.Controller {
		return accounts.NewGetAccountTokenController(h.db)
	})
}
