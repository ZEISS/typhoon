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
	"github.com/zeiss/typhoon/internal/web/controllers/users"
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
	return htmx.NewHxControllerHandler(login.NewIndexLoginController())
}

// Dashboard ...
func (h *handlers) Dashboard() fiber.Handler {
	return htmx.NewHxControllerHandler(dashboard.NewIndexDashboardController())
}

// Me ...
func (h *handlers) Me() fiber.Handler {
	return htmx.NewHxControllerHandler(me.NewMeController())
}

// ListOperators ...
func (h *handlers) ListOperators() fiber.Handler {
	return htmx.NewHxControllerHandler(operators.NewListOperatorsController(h.db))
}

// NewOperator ...
func (h *handlers) NewOperator() fiber.Handler {
	return htmx.NewHxControllerHandler(operators.NewOperatorController(h.db))
}

// CreateOperator ...
func (h *handlers) CreateOperator() fiber.Handler {
	return htmx.NewHxControllerHandler(operators.NewOperatorController(h.db))
}

// ShowOperator ...
func (h *handlers) ShowOperator() fiber.Handler {
	return htmx.NewHxControllerHandler(operators.NewShowOperatorController(h.db))
}

// DeleteOperator ...
func (h *handlers) DeleteOperator() fiber.Handler {
	return htmx.NewHxControllerHandler(operators.NewDeleteOperatorController(h.db))
}

// ListAccounts ...
func (h *handlers) ListAccounts() fiber.Handler {
	return htmx.NewHxControllerHandler(accounts.NewListAccountsController(h.db))
}

// NewAccount ...
func (h *handlers) NewAccount() fiber.Handler {
	return htmx.NewHxControllerHandler(accounts.NewAccountController(h.db))
}

// ListUsers ...
func (h *handlers) ListUsers() fiber.Handler {
	return htmx.NewHxControllerHandler(users.NewListUsersController(h.db))
}

// CreateAccount ...
func (h *handlers) CreateAccount() fiber.Handler {
	return htmx.NewHxControllerHandler(accounts.NewCreateController(h.db))
}

// ShowAccount ...
func (h *handlers) ShowAccount() fiber.Handler {
	return htmx.NewHxControllerHandler(accounts.NewShowAccountController(h.db))
}

// NewOperatorSkg ...
func (h *handlers) NewOperatorSkg() fiber.Handler {
	return htmx.NewHxControllerHandler(oskgs.NewSkgsController(h.db))
}

// CreateOperatorSkg ...
func (h *handlers) CreateOperatorSkg() fiber.Handler {
	return htmx.NewHxControllerHandler(oskgs.NewCreateSkgsController(h.db))
}

// OperatorSkgsOptions ...
func (h *handlers) OperatorSkgsOptions() fiber.Handler {
	return htmx.NewHxControllerHandler(pa.NewOperatorSkgsOptions(h.db))
}
