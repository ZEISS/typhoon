package search

import (
	"context"

	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/pkg/conv"
)

var _ = htmx.Controller(&SearchOperatorsControllerImpl{})

// Search ...
type SearchOperatorsControllerImpl struct {
	operators tables.Results[models.Operator]
	store     ports.Datastore
	htmx.DefaultController
}

// NewSearchOperatorsController ...
func NewSearchOperatorsController(store ports.Datastore) *SearchOperatorsControllerImpl {
	return &SearchOperatorsControllerImpl{
		operators: tables.Results[models.Operator]{SearchFields: []string{"name"}},
		store:     store,
	}
}

// Error ...
func (l *SearchOperatorsControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Prepare ...
func (l *SearchOperatorsControllerImpl) Prepare() error {
	var params struct {
		OperatorID string `json:"operator" form:"operator" query:"operator" validate:"required"`
	}

	err := l.BindQuery(&params)
	if err != nil {
		return err
	}
	l.operators.Search = params.OperatorID

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListOperators(ctx, &l.operators)
	})
}

// Get ...
func (l *SearchOperatorsControllerImpl) Get() error {
	return l.Render(
		htmx.Fragment(
			htmx.ForEach(l.operators.GetRows(), func(e *models.Operator, idx int) htmx.Node {
				return htmx.Option(
					htmx.Value(e.Name),
					htmx.Text(conv.String(e.ID)),
					htmx.Input(
						htmx.Type("hidden"),
						htmx.Name("operator_id"),
						htmx.Value(conv.String(e.ID)),
					),
				)
			})...,
		),
	)
}
