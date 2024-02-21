package transformation

import (
	"fmt"
	"strings"

	"github.com/zeiss/typhoon/pkg/apis/flow/v1alpha1"
	"github.com/zeiss/typhoon/pkg/flow/adapter/transformation/common/storage"
	"github.com/zeiss/typhoon/pkg/flow/adapter/transformation/transformer"
	"github.com/zeiss/typhoon/pkg/flow/adapter/transformation/transformer/add"
	"github.com/zeiss/typhoon/pkg/flow/adapter/transformation/transformer/delete"
	"github.com/zeiss/typhoon/pkg/flow/adapter/transformation/transformer/parse"
	"github.com/zeiss/typhoon/pkg/flow/adapter/transformation/transformer/shift"
	"github.com/zeiss/typhoon/pkg/flow/adapter/transformation/transformer/store"
)

const (
	defaultEventPathSeparator = "."
)

// Pipeline is a set of Transformations that are
// sequentially applied to JSON data.
type Pipeline struct {
	Transformers []transformer.Transformer
	Storage      *storage.Storage
}

// register loads available Transformation into a named map.
func register() map[string]transformer.Transformer {
	transformations := make(map[string]transformer.Transformer)

	add.Register(transformations)
	delete.Register(transformations)
	shift.Register(transformations)
	store.Register(transformations)
	parse.Register(transformations)

	return transformations
}

// newPipeline loads available Transformations and creates a Pipeline.
func newPipeline(transformations []v1alpha1.Transform, storage *storage.Storage) (*Pipeline, error) {
	availableTransformers := register()
	pipeline := []transformer.Transformer{}

	for _, transformation := range transformations {
		operation, exist := availableTransformers[transformation.Operation]
		if !exist {
			return nil, fmt.Errorf("transformation %q not found", transformation.Operation)
		}
		for _, kv := range transformation.Paths {
			separator := defaultEventPathSeparator
			if kv.Separator != "" {
				separator = kv.Separator
			}
			transformer := operation.New(kv.Key, kv.Value, separator)
			transformer.SetStorage(storage)
			pipeline = append(pipeline, transformer)
		}
	}

	return &Pipeline{
		Transformers: pipeline,
		Storage:      storage,
	}, nil
}

// Apply applies Pipeline transformations.
func (p *Pipeline) apply(eventID string, data []byte, init bool) ([]byte, error) {
	var err error
	var errs []string
	for _, v := range p.Transformers {
		if init == v.InitStep() {
			if data, err = v.Apply(eventID, data); err != nil {
				errs = append(errs, err.Error())
			}
		}
	}
	if len(errs) != 0 {
		return data, fmt.Errorf(strings.Join(errs, ","))
	}
	return data, nil
}
