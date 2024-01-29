package adapter

import (
	openapi "github.com/zeiss/typhoon/api"
	"github.com/zeiss/typhoon/internal/adapter/handlers"
)

var _ openapi.StrictServerInterface = (*Handlers)(nil)

// Handlers ...
type Handlers struct {
	*handlers.TeamsHandler
	*handlers.SystemsHandler
	*handlers.VersionHandler
}

// NewHandlers ...
func NewHandlers(teams *handlers.TeamsHandler, systems *handlers.SystemsHandler, version *handlers.VersionHandler) *Handlers {
	return &Handlers{teams, systems, version}
}
