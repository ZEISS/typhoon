package sysaccount

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nats-io/jwt"
	"github.com/nats-io/nkeys"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/utils"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// UpdateSystemAccountControllerParams ...
type UpdateSystemAccountControllerParams struct {
	ID uuid.UUID `json:"id" params:"id" validate:"required:uuid"`
}

// UpdateSystemAccountControllerBody ...
type UpdateSystemAccountControllerBody struct {
	SystemAccountID uuid.UUID `json:"system_account_id" form:"system_account_id" validate:"required,uuid"`
}

// UpdateSystemAccountControllerImpl ...
type UpdateSystemAccountControllerImpl struct {
	Params UpdateSystemAccountControllerParams
	Body   UpdateSystemAccountControllerBody

	store ports.Datastore
	htmx.DefaultController
}

// NewCreateSkgsController ...
func NewCreateSkgsController(store ports.Datastore) *UpdateSystemAccountControllerImpl {
	return &UpdateSystemAccountControllerImpl{
		Params:            UpdateSystemAccountControllerParams{},
		Body:              UpdateSystemAccountControllerBody{},
		store:             store,
		DefaultController: htmx.DefaultController{},
	}
}

// Prepare ...
func (l *UpdateSystemAccountControllerImpl) Prepare() error {
	err := l.BindParams(&l.Params)
	if err != nil {
		return err
	}

	err = l.BindBody(&l.Body)
	if err != nil {
		return err
	}

	return nil
}

// Put ...
func (l *UpdateSystemAccountControllerImpl) Put() error {
	op := models.Operator{ID: l.Params.ID}
	acc := models.Account{ID: l.Body.SystemAccountID}

	err := l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		// TODO: implement a parallel resolution
		err := tx.GetOperator(l.Context(), &op)
		if err != nil {
			return err
		}

		err = tx.GetAccount(l.Context(), &acc)
		if err != nil {
			return err
		}
		op.SystemAdminAccountID = utils.PtrUUID(acc.ID)

		return nil
	})

	pk, err := nkeys.FromSeed(op.Key.Seed)
	if err != nil {
		return err
	}

	id, err := pk.PublicKey()
	if err != nil {
		return err
	}

	oc := jwt.NewOperatorClaims(id)
	oc.Name = op.Name
	oc.SystemAccount = acc.Key.ID

	for _, sk := range op.SigningKeyGroups {
		oc.SigningKeys.Add(sk.Key.ID, sk.Key.ID, sk.Key.ID)
	}

	token, err := oc.Encode(pk)
	if err != nil {
		return err
	}
	op.Token = models.Token{
		ID:    id,
		Token: token,
	}

	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.UpdateOperator(l.Context(), &op)
	})
	if err != nil {
		return err
	}

	accs := []*models.Account{}
	for _, acc := range op.Accounts {
		accs = append(accs, &acc)
	}

	return htmx.RenderComp(
		l.Ctx(),
		htmx.Fragment(
			htmx.Div(
				htmx.ID("alerts"),
				htmx.HxSwapOob("true"),
				toasts.ToastEnd(
					toasts.ToastProps{},
					// toasts.ToastAlertInfo(
					// 	htmx.Text("Info"),
					// ),
					// toasts.ToastAlertError(
					// 	htmx.Text("Error"),
					// ),
					toasts.ToastAlertSuccess(
						icons.CheckCircleOutline(
							icons.IconProps{},
						),
						htmx.Span(htmx.Text("Success")),
					),
				),
			),
			forms.SelectBordered(
				forms.SelectProps{},
				htmx.HxPut(fmt.Sprintf("/operators/%s/system-account", op.ID)),
				htmx.HxTarget("this"),
				forms.Option(
					forms.OptionProps{
						Selected: true,
						Disabled: true,
					},
					htmx.Text("Select account"),
				),
				htmx.Name("system_account_id"),
				htmx.Group(
					htmx.ForEach(accs, func(account *models.Account, idx int) htmx.Node {
						return forms.Option(
							forms.OptionProps{
								Value:    account.ID.String(),
								Selected: op.SystemAdminAccountID != nil && account.ID == utils.UUIDPtr(op.SystemAdminAccountID),
							},
							htmx.Text(account.Name),
						)
					})...,
				),
			),
		),
	)
}
