package resolvers

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

const (
	// ValuesKeyOperators ...
	ValuesKeyOperators = "user"
)

// ListOperators ...
func ListOperators(db ports.Repository) htmx.ResolveFunc {
	return func(ctx *fiber.Ctx) (interface{}, interface{}, error) {
		query := struct {
			Limit  int    `json:"limit" xml:"limit" form:"limit"`
			Offset int    `json:"offset" xml:"offset" form:"offset"`
			Search string `json:"search" xml:"search" form:"search"`
		}{}

		err := ctx.QueryParser(&query)
		if err != nil {
			return nil, nil, err
		}

		ops := models.Pagination[models.Operator]{}

		ops.Limit = query.Limit
		ops.Offset = query.Offset
		ops.Search = query.Search

		err = db.ListOperators(ctx.Context(), &ops)
		if err != nil {
			return nil, nil, err
		}

		return ValuesKeyOperators, ops, nil
	}
}
