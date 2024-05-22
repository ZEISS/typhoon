package handlers

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/web/controllers/dashboard"
	"github.com/zeiss/typhoon/internal/web/controllers/login"
	"github.com/zeiss/typhoon/internal/web/controllers/me"
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