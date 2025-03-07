package v1alpha1

import (
	"context"
)

const (
	defaultSchedule = "* * * * *"
)

// SetDefaults sets the default values for the PingSource.
func (s *PingSource) SetDefaults(ctx context.Context) {
	s.Spec.SetDefaults(ctx)
}

// SetDefaults sets the default values for the PingSourceSpec.
func (ss *PingSourceSpec) SetDefaults(ctx context.Context) {
	if ss.Schedule == "" {
		ss.Schedule = defaultSchedule
	}
}
