package transformer

import (
	"github.com/zeiss/typhoon/pkg/flow/adapter/transformation/common/storage"
)

// Transformer is an interface that contains common methods
// to work with JSON data.
type Transformer interface {
	New(key, value, separator string) Transformer
	Apply(eventID string, data []byte) ([]byte, error)
	SetStorage(*storage.Storage)
	InitStep() bool
}
