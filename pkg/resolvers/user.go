package resolvers

import (
	"github.com/zeiss/typhoon/internal/web/ports"

	"github.com/gofiber/fiber/v2"
	goth "github.com/zeiss/fiber-goth"
	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
)

const (
	// ValuesKeyUser ...
	ValuesKeyUser = "user"
)

// UserByID ...
func UserByID(db ports.Repository) htmx.ResolveFunc {
	return func(ctx *fiber.Ctx) (interface{}, interface{}, error) {
		session, err := goth.SessionFromContext(ctx)
		if err != nil {
			return nil, nil, err
		}

		user := adapters.GothUser{
			ID: session.UserID,
		}

		err = db.GetUser(ctx.Context(), &user)
		if err != nil {
			return err, nil, err
		}

		return ValuesKeyUser, user, nil
	}
}
