package filter

import (
	"sync"

	"github.com/zeiss/typhoon/pkg/routing/eventfilter/cel"
	"k8s.io/apimachinery/pkg/types"
)

type (
	filterGenerations map[int64]cel.ConditionalFilter
	filterUIDs        map[types.UID]filterGenerations
)

type expressionStorage struct {
	*sync.RWMutex
	filterUIDs
}

func newExpressionStorage() *expressionStorage {
	return &expressionStorage{
		RWMutex:    &sync.RWMutex{},
		filterUIDs: make(filterUIDs),
	}
}

func (f *expressionStorage) get(uid types.UID, generation int64) (cel.ConditionalFilter, bool) {
	f.RLock()
	defer f.RUnlock()

	filterGens, exist := f.filterUIDs[uid]
	if !exist {
		return cel.ConditionalFilter{}, false
	}

	filter, exist := filterGens[generation]
	return filter, exist
}

// set method overrides previous generations of compiled expressions
func (f *expressionStorage) set(uid types.UID, generation int64, condition cel.ConditionalFilter) {
	f.Lock()
	defer f.Unlock()

	f.filterUIDs[uid] = filterGenerations{
		generation: condition,
	}
}
