package accounts

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/nats-io/jwt/v2"
	"github.com/nats-io/nkeys"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/utils"
	"github.com/zeiss/typhoon/internal/web/components/alerts"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var validate *validator.Validate

// CreateControllerBody ...
type CreateControllerBody struct {
	OperatorID                  uuid.UUID `json:"operator_id" form:"operator_id" validate:"required,uuid"`
	Name                        string    `json:"name" form:"name" validate:"required,min=3,max=100"`
	Description                 string    `json:"description" form:"description" validate:"required,min=3,max=1024"`
	JetStreamEnable             bool      `json:"jetstream_enable" form:"jetstream_enable"`
	JetStreamMaxDiskStorage     float64   `json:"jetstream_max_disk_storage" form:"jetstream_max_disk_storage"`
	JetStreamMaxDiskStorageUnit string    `json:"jetstream_max_disk_storage_unit" form:"jetstream_max_disk_storage_unit"`
	JetStreamMaxStreams         int64     `json:"jetstream_max_streams" form:"jetstream_max_streams"`
	JetStreamMaxConsumers       int64     `json:"jetstream_max_consumers" form:"jetstream_max_consumers"`
	JetStreamMaxStreamSize      float64   `json:"jetstream_max_stream_size" form:"jetstream_max_stream_size"`
	JetStreamMaxStreamSizeUnit  string    `json:"jetstream_max_stream_size_unit" form:"jetstream_max_stream_size_unit"`
	JetStreamMaxBytesRequired   bool      `json:"jetstream_max_bytes_required" form:"jetstream_max_bytes_required"`
}

// CreateControllerImpl ...
type CreateControllerImpl struct {
	Form CreateControllerBody

	ports.Repository
	htmx.DefaultController
}

// NewCreateController ...
func NewCreateController(db ports.Repository) *CreateControllerImpl {
	return &CreateControllerImpl{
		Repository:        db,
		DefaultController: htmx.DefaultController{},
	}
}

// Prepare ...
func (l *CreateControllerImpl) Prepare() error {
	validate = validator.New()

	err := l.Ctx().BodyParser(&l.Form)
	if err != nil {
		return err
	}

	err = validate.Struct(l)
	if err != nil {
		return err
	}

	return nil
}

// Error ...
func (l *CreateControllerImpl) Error(err error) error {
	return htmx.RenderComp(
		l.Ctx(),
		alerts.Error(
			alerts.ErrorProps{
				Error: err,
				ID:    "alerts",
			},
		),
		htmx.RenderStatusCode(err),
	)
}

// Post ...
func (l *CreateControllerImpl) Post() error {
	account := models.Account{
		Name:        l.Form.Name,
		OperatorID:  l.Form.OperatorID,
		Description: utils.StrPtr(l.Form.Description),
	}

	operator := models.Operator{
		ID: l.Form.OperatorID,
	}
	err := l.GetOperator(l.Context(), &operator)
	if err != nil {
		return err
	}

	pk, err := nkeys.CreateAccount()
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
	account.Key = models.NKey{ID: id, Seed: seed}

	skg := models.SigningKeyGroup{Name: "Default", Description: "Default signing key group"}

	skgpk, err := nkeys.CreateAccount()
	if err != nil {
		return err
	}

	skgid, err := skgpk.PublicKey()
	if err != nil {
		return err
	}

	skgseed, err := skgpk.Seed()
	if err != nil {
		return err
	}
	skg.Key = models.NKey{ID: skgid, Seed: skgseed}
	account.SigningKeyGroups = append(account.SigningKeyGroups, skg)

	// @katallaxie: this is a bit weird, but I think it's a good idea to have a default signing key group
	osk, err := nkeys.FromSeed(operator.SigningKeyGroups[0].Key.Seed)
	if err != nil {
		return err
	}

	ac := jwt.NewAccountClaims(id)
	ac.Name = l.Form.Name
	ac.Issuer = operator.Key.ID
	ac.SigningKeys.Add(skg.Key.ID)

	token, err := ac.Encode(osk)
	if err != nil {
		return err
	}
	account.Token = models.Token{ID: id, Token: token}

	err = l.CreateAccount(l.Context(), &account)
	if err != nil {
		return err
	}

	htmx.Redirect(l.Ctx(), "/accounts")

	return nil
}
