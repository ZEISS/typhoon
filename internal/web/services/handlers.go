package services

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/web/controllers/dashboard"
	"github.com/zeiss/typhoon/internal/web/controllers/login"
	"github.com/zeiss/typhoon/internal/web/controllers/me"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var _ ports.Handlers = (*handlers)(nil)

type handlers struct{}

// NewHandlers ...
func NewHandlers() *handlers {
	return &handlers{}
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
