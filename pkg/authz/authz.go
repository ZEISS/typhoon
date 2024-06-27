package authz

import (
	"github.com/gofiber/fiber/v2"
	authz "github.com/zeiss/fiber-authz"
)

var _ authz.AuthzPrincipalResolver = (*AuthzPrincipalResolver)(nil)
var _ authz.AuthzObjectResolver = (*AuthzObjectResolver)(nil)
var _ authz.AuthzActionResolver = (*AuthzActionResolver)(nil)

// AuthzPrincipalResolver ...
type AuthzPrincipalResolver struct {
}

// Resolve ...
func (a *AuthzPrincipalResolver) Resolve(c *fiber.Ctx) (authz.AuthzPrincipal, error) {
	return authz.AuthzPrincipal(""), nil
}

// AuthzObjectResolver ...
type AuthzObjectResolver struct {
}

// Resolve ...
func (a *AuthzObjectResolver) Resolve(c *fiber.Ctx) (authz.AuthzObject, error) {
	return authz.AuthzObject(""), nil
}

// AuthzActionResolver ...
type AuthzActionResolver struct {
}

// Resolve ...
func (a *AuthzActionResolver) Resolve(c *fiber.Ctx) (authz.AuthzAction, error) {
	return authz.AuthzAction(""), nil
}
