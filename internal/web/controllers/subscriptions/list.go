package subscriptions

import (
	"context"

	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/subscriptions"
	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tables"
)

var _ = htmx.Controller(&ListSubscriptionsController{})

// ListSubscriptionsController ...
type ListSubscriptionsController struct {
	subscriptions tables.Results[models.Subscription]
	store         ports.Datastore
	htmx.DefaultController
}

// NewListSubscriptionsController ...
func NewListSubscriptionsController(store ports.Datastore) *ListSubscriptionsController {
	return &ListSubscriptionsController{store: store}
}

// Prepare ...
func (l *ListSubscriptionsController) Prepare() error {
	err := l.BindQuery(&l.subscriptions)
	if err != nil {
		return err
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListSubscriptions(ctx, &l.subscriptions)
	})
}

// Get ...
func (l *ListSubscriptionsController) Get() error {
	return l.Render( // render the html using htmx
		components.Page(
			components.PageProps{
				Title: "Subscriptions",
				Boost: true,
			},
			components.Layout(
				components.LayoutProps{
					Path: l.Path(), // get the current path
				},
				cards.CardBordered(
					cards.CardProps{},
					cards.Body(
						cards.BodyProps{},
						subscriptions.SubscriptionsTable(
							subscriptions.SubscriptionsTableProps{
								Limit:         l.subscriptions.GetLimit(),
								Offset:        l.subscriptions.GetOffset(),
								Total:         l.subscriptions.GetLen(),
								Subscriptions: l.subscriptions.GetRows(),
							},
						),
						cards.Actions(
							cards.ActionsProps{},
							htmx.A(
								htmx.Href("subscriptions/new"),
								buttons.Outline(
									buttons.ButtonProps{
										ClassNames: htmx.ClassNames{
											"btn-sm": true,
										},
									},
									htmx.Text("Create Subscription"),
								),
							),
						),
					),
				),
			),
		),
	)
}
