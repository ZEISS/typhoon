package accounts

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/nats-io/jwt/v2"
	"github.com/nats-io/nkeys"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/utils"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var validate *validator.Validate

// CreateControllerBody ...
type CreateControllerBody struct {
	OperatorID                  uuid.UUID `json:"operator_id" form:"operator_id" validate:"required,uuid"`
	OperatorSigningKeyGroupID   string    `json:"operator_signing_key_group_id" form:"operator_skgs_id"`
	TeamID                      uuid.UUID `json:"team_id" form:"team_id" validate:"required,uuid"`
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

	store ports.Datastore
	htmx.TransactionController
	htmx.DefaultController
}

// NewCreateController ...
func NewCreateController(store ports.Datastore) *CreateControllerImpl {
	return &CreateControllerImpl{store: store}
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
	return toasts.RenderToasts(
		l.Ctx(),
		toasts.Toasts(
			toasts.ToastsProps{},
			toasts.ToastAlertError(
				toasts.ToastProps{},
				htmx.Text(err.Error()),
			),
		),
	)
}

// Post ...
func (l *CreateControllerImpl) Post() error {
	account := models.Account{
		Name:                           l.Form.Name,
		Description:                    utils.StrPtr(l.Form.Description),
		LimitJetStreamMaxDiskStorage:   utils.PrettyByteSize(l.Form.JetStreamMaxDiskStorage, l.Form.JetStreamMaxDiskStorageUnit),
		LimitJetStreamMaxStreams:       l.Form.JetStreamMaxStreams,
		LimitJetStreamMaxConsumers:     l.Form.JetStreamMaxConsumers,
		LimitJetStreamMaxStreamBytes:   utils.PrettyByteSize(l.Form.JetStreamMaxStreamSize, l.Form.JetStreamMaxStreamSizeUnit),
		LimitJetStreamMaxBytesRequired: l.Form.JetStreamMaxBytesRequired,
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

	osgk := models.NKey{ID: string(l.Form.OperatorSigningKeyGroupID)}

	err = l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetNKey(ctx, &osgk)
	})
	if err != nil {
		return err
	}

	osk, err := nkeys.FromSeed(osgk.Seed)
	if err != nil {
		return err
	}

	ac := jwt.NewAccountClaims(id)
	ac.Name = l.Form.Name
	ac.Issuer = osgk.ID
	ac.SigningKeys.Add(skg.Key.ID)
	ac.Limits.JetStreamLimits.DiskStorage = -1
	ac.Limits.JetStreamLimits.Streams = -1

	token, err := ac.Encode(osk)
	if err != nil {
		return err
	}
	account.Token = models.Token{ID: id, Token: token}
	account.Signer = osgk
	account.OwnerType = models.TeamAccount
	account.OwnerID = l.Form.TeamID

	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateAccount(ctx, &account)
	})
	if err != nil {
		return err
	}

	htmx.Redirect(l.Ctx(), "/accounts")

	return nil
}
