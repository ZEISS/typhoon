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
	// NewOperator ...
	NewOperator() fiber.Handler
	// CreateOperator ...
	CreateOperator() fiber.Handler
	// ShowOperator ...
	ShowOperator() fiber.Handler
	// DeleteOperator ...
	DeleteOperator() fiber.Handler
	// ListAccounts ...
	ListAccounts() fiber.Handler
	// NewAccount ...
	NewAccount() fiber.Handler
	// CreateAccount ...
	CreateAccount() fiber.Handler
	// ShowAccount ...
	ShowAccount() fiber.Handler
	// // DeleteAccount ...
	// DeleteAccount() fiber.Handler
	// ListUsers ...
	ListUsers() fiber.Handler
	// NewOperatorSkg ...
	NewOperatorSkg() fiber.Handler
	// CreateOperatorSkg ...
	CreateOperatorSkg() fiber.Handler
	// OperatorSkgsOptions ...
	OperatorSkgsOptions() fiber.Handler
	// AccountSksOptions ...
	AccountSksOptions() fiber.Handler
	// NewUser ...
	NewUser() fiber.Handler
}
