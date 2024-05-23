package handlers

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/web/controllers/dashboard"
	"github.com/zeiss/typhoon/internal/web/controllers/login"
	"github.com/zeiss/typhoon/internal/web/controllers/me"
	"github.com/zeiss/typhoon/internal/web/controllers/operators"
	"github.com/zeiss/typhoon/internal/web/ports"
	"github.com/zeiss/typhoon/pkg/resolvers"
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
	return htmx.NewHxControllerHandler(me.NewMeController(), htmx.Config{Resolvers: []htmx.ResolveFunc{resolvers.UserByID(h.db)}})
}

// ListOperators ...
func (h *handlers) ListOperators() fiber.Handler {
	return htmx.NewHxControllerHandler(operators.NewListOperatorsController(h.db), htmx.Config{Resolvers: []htmx.ResolveFunc{resolvers.ListOperators(h.db)}})
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
