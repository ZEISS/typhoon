package operators

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/nats-io/jwt/v2"
	"github.com/nats-io/nkeys"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/components/alerts"
	"github.com/zeiss/typhoon/internal/web/ports"
)

const (
	operatorsShowURL = "/operators/%s"
)

var validate *validator.Validate

type createControllerBody struct {
	Name        string `json:"name" form:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" form:"description" validate:"required,min=3,max=1024"`
}

// CreateControllerImpl ...
type CreateControllerImpl struct {
	body  createControllerBody
	store ports.Datastore
	htmx.DefaultController
}

// NewCreateController ...
func NewCreateController(store ports.Datastore) *CreateControllerImpl {
	return &CreateControllerImpl{store: store}
}

// Prepare ...
func (l *CreateControllerImpl) Prepare() error {
	validate = validator.New()

	err := l.Ctx().BodyParser(&l.body)
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
	return l.Render(
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
	operator, err := models.NewOperator(l.body.Name, l.body.Description)
	if err != nil {
		return err
	}

	// Create operator signing key group
	oskg := models.SigningKeyGroup{Name: "Default", Description: "Default signing key group"}
	opk, err := nkeys.CreateOperator()
	if err != nil {
		return err
	}

	id, err := opk.PublicKey()
	if err != nil {
		return err
	}

	oseed, err := opk.Seed()
	if err != nil {
		return err
	}
	oskg.Key = models.NKey{ID: id, Seed: oseed}
	operator.SigningKeyGroups = append(operator.SigningKeyGroups, oskg)

	// Setup account
	account := models.Account{
		Name: "System Account",
	}

	apk, err := nkeys.CreateAccount()
	if err != nil {
		return err
	}

	aid, err := apk.PublicKey()
	if err != nil {
		return err
	}

	aseed, err := apk.Seed()
	if err != nil {
		return err
	}
	account.Key = models.NKey{ID: aid, Seed: aseed}

	askg := models.SigningKeyGroup{Name: "Default", Description: "Default signing key group"}
	askgpk, err := nkeys.CreateAccount()
	if err != nil {
		return err
	}

	askgid, err := askgpk.PublicKey()
	if err != nil {
		return err
	}

	skgseed, err := askgpk.Seed()
	if err != nil {
		return err
	}
	askg.Key = models.NKey{ID: askgid, Seed: skgseed}
	account.SigningKeyGroups = append(account.SigningKeyGroups, askg)

	// Create account claim
	ac := jwt.NewAccountClaims(aid)
	ac.Name = "System Account"
	ac.Issuer = operator.Key.ID
	ac.SigningKeys.Add(askg.Key.ID)

	ac.Exports = jwt.Exports{&jwt.Export{
		Name:                 "account-monitoring-services",
		Subject:              "$SYS.REQ.ACCOUNT.*.*",
		Type:                 jwt.Service,
		ResponseType:         jwt.ResponseTypeStream,
		AccountTokenPosition: 4,
		Info: jwt.Info{
			Description: `Request account specific monitoring services for: SUBSZ, CONNZ, LEAFZ, JSZ and INFO`,
			InfoURL:     "https://docs.nats.io/nats-server/configuration/sys_accounts",
		},
	}, &jwt.Export{
		Name:                 "account-monitoring-streams",
		Subject:              "$SYS.ACCOUNT.*.>",
		Type:                 jwt.Stream,
		AccountTokenPosition: 3,
		Info: jwt.Info{
			Description: `Account specific monitoring stream`,
			InfoURL:     "https://docs.nats.io/nats-server/configuration/sys_accounts",
		},
	}}

	token, err := ac.Encode(opk)
	if err != nil {
		return err
	}
	account.Token = models.Token{ID: aid, Token: token}

	// Create operator claim
	oc := jwt.NewOperatorClaims(id)
	oc.Name = operator.Name

	for _, sk := range operator.SigningKeyGroups {
		oc.SigningKeys.Add(sk.Key.ID, sk.Key.ID, sk.Key.ID)
	}
	oc.SystemAccount = account.Key.ID

	token, err = oc.Encode(opk)
	if err != nil {
		return err
	}
	operator.Token = models.Token{ID: id, Token: token}

	// Associate account with operator
	operator.SystemAccount = account

	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateOperator(ctx, &operator)
	})
	if err != nil {
		return err
	}

	return l.Redirect(fmt.Sprintf(operatorsShowURL, operator.ID))
}
