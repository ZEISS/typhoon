package v1alpha1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

func TestCloudEventsSourceGetStatus(t *testing.T) {
	r := &CloudEventsSource{
		Status: v1alpha1.Status{},
	}
	if got, want := r.GetStatus(), &r.Status.Status; got != want {
		t.Errorf("GetStatus=%v, want=%v", got, want)
	}
}

func TestCloudEventsSourceGetSink(t *testing.T) {
	d := duckv1.Destination{
		Ref: &duckv1.KReference{
			Kind:       "TestKind",
			APIVersion: "v1alpha1",
			Namespace:  "testnamespace",
			Name:       "testname",
		},
	}

	r := &CloudEventsSource{
		Spec: CloudEventsSourceSpec{
			SourceSpec: duckv1.SourceSpec{
				Sink: d,
			},
		},
	}
	if got, want := *r.GetSink(), d; got != want {
		t.Errorf("GetSink=%v, want=%v", got, want)
	}
}

func TestCloudEventsSourceGetStatusManager(t *testing.T) {
	s := v1alpha1.Status{}
	sm := v1alpha1.StatusManager{
		ConditionSet: v1alpha1.DefaultConditionSet,
		Status:       &s,
	}
	r := &CloudEventsSource{
		Status: s,
	}

	assert.Equal(t, sm, *r.GetStatusManager(), "unexpected Status().ConditionSet")
}

func TestCloudEventsSourceGetGroupVersionKind(t *testing.T) {
	s := CloudEventsSource{}
	gvk := s.GetGroupVersionKind()
	if gvk.Kind != "CloudEventsSource" {
		t.Errorf("Should be CloudEventsSource.")
	}
}
