package skgs

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/nats-io/jwt"
	"github.com/nats-io/nkeys"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// CreateSkgsControllerImpl ...
type CreateSkgsControllerImpl struct {
	ID          uuid.UUID `json:"id" params:"id" validate:"required:uuid"`
	Name        string    `json:"name" form:"name" validate:"required"`
	Description string    `json:"description" form:"description" validate:"required"`

	ports.Operators
	htmx.DefaultController
}

// NewCreateSkgsController ...
func NewCreateSkgsController(db ports.Operators) *CreateSkgsControllerImpl {
	return &CreateSkgsControllerImpl{
		Operators:         db,
		DefaultController: htmx.DefaultController{},
	}
}

// Prepare ...
func (l *CreateSkgsControllerImpl) Prepare() error {
	err := l.BindParams(l)
	if err != nil {
		return err
	}

	err = l.BindBody(l)
	if err != nil {
		return err
	}

	return nil
}

// Post ...
func (l *CreateSkgsControllerImpl) Post() error {
	op := models.Operator{ID: l.ID}
	skg := models.SigningKeyGroup{Name: l.Name, Description: l.Description}

	err := l.GetOperator(l.Context(), &op)
	if err != nil {
		return err
	}

	pk, err := nkeys.CreateOperator()
	if err != nil {
		return err
	}

	id, err := pk.PublicKey()
	if err != nil {
		return err
	}

	seed, err := pk.Seed()
	if err != nil {
		return err
	}
	skg.Key = models.NKey{ID: id, Seed: seed}

	op.SigningKeyGroups = append(op.SigningKeyGroups, skg)

	oc := jwt.NewOperatorClaims(id)
	oc.Name = op.Name

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

	err = l.UpdateOperator(l.Context(), &op)
	if err != nil {
		return err
	}

	htmx.Redirect(l.Ctx(), fmt.Sprintf("/operators/%s", l.ID))

	return nil
}
