package delete

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/zeiss/typhoon/pkg/flow/adapter/transformation/common/storage"
	"github.com/zeiss/typhoon/pkg/flow/adapter/transformation/transformer"
)

var _ transformer.Transformer = (*Delete)(nil)

// Delete object implements Transformer interface.
type Delete struct {
	Path      string
	Value     string
	Type      string
	Separator string

	variables *storage.Storage
}

// InitStep is used to figure out if this operation should
// run before main Transformations. For example, Store
// operation needs to run first to load all Pipeline variables.
var InitStep bool = false

// operationName is used to identify this transformation.
var operationName string = "delete"

// Register adds this transformation to the map which will
// be used to create Transformation pipeline.
func Register(m map[string]transformer.Transformer) {
	m[operationName] = &Delete{}
}

// SetStorage sets a shared Storage with Pipeline variables.
func (d *Delete) SetStorage(storage *storage.Storage) {
	d.variables = storage
}

// InitStep returns "true" if this Transformation should run
// as init step.
func (d *Delete) InitStep() bool {
	return InitStep
}

// New returns a new instance of Delete object.
func (d *Delete) New(key, value, separator string) transformer.Transformer {
	return &Delete{
		Path:      key,
		Value:     value,
		Separator: separator,

		variables: d.variables,
	}
}

// Apply is a main method of Transformation that removed any type of
// variables from existing JSON.
func (d *Delete) Apply(eventID string, data []byte) ([]byte, error) {
	d.Value = d.retrieveString(eventID, d.Value)

	result, err := d.parse(data, "", "")
	if err != nil {
		return data, err
	}

	output, err := json.Marshal(result)
	if err != nil {
		return data, err
	}

	return output, nil
}

func (d *Delete) retrieveString(eventID, key string) string {
	if value := d.variables.Get(eventID, key); value != nil {
		if str, ok := value.(string); ok {
			return str
		}
	}
	return key
}

// nolint:gocyclo
func (d *Delete) parse(data interface{}, key, path string) (interface{}, error) { // nolint:unparam
	output := make(map[string]interface{})
	// TODO: keep only one filter call
	if d.filter(path, data) {
		return nil, nil
	}

	switch value := data.(type) {
	case []byte:
		var m interface{}
		if err := json.Unmarshal(value, &m); err != nil {
			return nil, fmt.Errorf("unmarshal err: %w", err)
		}
		o, err := d.parse(m, key, path)
		if err != nil {
			return nil, fmt.Errorf("recursive call in []bytes case: %w", err)
		}
		return o, nil
	case float64, bool, string, nil:
		return value, nil
	case []interface{}:
		slice := []interface{}{}
		for i, v := range value {
			o, err := d.parse(v, key, fmt.Sprintf("%s[%d]", path, i))
			if err != nil {
				return nil, fmt.Errorf("recursive call in []interface case: %w", err)
			}
			slice = append(slice, o)
		}
		return slice, nil
	case map[string]interface{}:
		for k, v := range value {
			subPath := fmt.Sprintf("%s.%s", path, k)
			if d.filter(subPath, v) {
				continue
			}
			o, err := d.parse(v, k, subPath)
			if err != nil {
				return nil, fmt.Errorf("recursive call in map[]interface case: %w", err)
			}
			output[k] = o
		}
	}

	return output, nil
}

func (d *Delete) filter(path string, value interface{}) bool {
	switch {
	case d.Path != "" && d.Value != "":
		return d.filterPathAndValue(path, value)
	case d.Path != "":
		return d.filterPath(path)
	case d.Value != "":
		return d.filterValue(value)
	}
	// consider empty key and path as "delete any"
	return true
}

func (d *Delete) filterPath(path string) bool {
	return d.Separator+d.Path == path
}

func (d *Delete) filterValue(value interface{}) bool {
	switch v := value.(type) {
	case string:
		return v == d.Value
	case float64:
		return d.Value == strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return d.Value == fmt.Sprintf("%t", v)
	}
	return false
}

func (d *Delete) filterPathAndValue(path string, value interface{}) bool {
	return d.filterPath(path) && d.filterValue(value)
}
