package store

import (
	"encoding/json"
	"strings"

	"github.com/zeiss/typhoon/pkg/flow/adapter/transformation/common"
	"github.com/zeiss/typhoon/pkg/flow/adapter/transformation/common/convert"
	"github.com/zeiss/typhoon/pkg/flow/adapter/transformation/common/storage"
	"github.com/zeiss/typhoon/pkg/flow/adapter/transformation/transformer"
)

var _ transformer.Transformer = (*Store)(nil)

// Store object implements Transformer interface.
type Store struct {
	Path      string
	Value     string
	Separator string

	variables *storage.Storage
}

// InitStep is used to figure out if this operation should
// run before main Transformations. For example, Store
// operation needs to run first to load all Pipeline variables.
var InitStep bool = true

// operationName is used to identify this transformation.
var operationName string = "store"

// Register adds this transformation to the map which will
// be used to create Transformation pipeline.
func Register(m map[string]transformer.Transformer) {
	m[operationName] = &Store{}
}

// SetStorage sets a shared Storage with Pipeline variables.
func (s *Store) SetStorage(storage *storage.Storage) {
	s.variables = storage
}

// InitStep returns "true" if this Transformation should run
// as init step.
func (s *Store) InitStep() bool {
	return InitStep
}

// New returns a new instance of Store object.
func (s *Store) New(key, value, separator string) transformer.Transformer {
	return &Store{
		Path:      key,
		Value:     value,
		Separator: separator,

		variables: s.variables,
	}
}

// Apply is a main method of Transformation that stores JSON values
// into variables that can be used by other Transformations in a pipeline.
func (s *Store) Apply(eventID string, data []byte) ([]byte, error) {
	path := convert.SliceToMap(strings.Split(s.Value, s.Separator), "")

	var event interface{}
	if err := json.Unmarshal(data, &event); err != nil {
		return data, err
	}

	value := common.ReadValue(event, path)

	s.variables.Set(eventID, s.Path, value)

	return data, nil
}
