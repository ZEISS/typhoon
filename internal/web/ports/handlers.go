package ports

import "github.com/gofiber/fiber/v2"

// Handlers ...
type Handlers interface {
	// Login ...
	Login() fiber.Handler
	// Dashboard ...
	Dashboard() fiber.Handler
	// Me ...
	Me() fiber.Handler
	// ListOperators ...
	ListOperators() fiber.Handler
}
