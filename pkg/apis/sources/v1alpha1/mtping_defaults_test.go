package v1alpha1

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPingSourceSetDefaults(t *testing.T) {
	testCases := map[string]struct {
		initial  PingSource
		expected PingSource
	}{
		"nil": {
			expected: PingSource{
				Spec: PingSourceSpec{
					Schedule: defaultSchedule,
				},
			},
		},
		"empty": {
			initial: PingSource{},
			expected: PingSource{
				Spec: PingSourceSpec{
					Schedule: defaultSchedule,
				},
			},
		},
		"with schedule": {
			initial: PingSource{
				Spec: PingSourceSpec{
					Schedule: "1 2 3 4 5",
				},
			},
			expected: PingSource{
				Spec: PingSourceSpec{
					Schedule: "1 2 3 4 5",
				},
			},
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			tc.initial.SetDefaults(context.TODO())
			if diff := cmp.Diff(tc.expected, tc.initial); diff != "" {
				t.Fatal("Unexpected defaults (-want, +got):", diff)
			}
		})
	}
}
