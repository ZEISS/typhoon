// Package structs provides helpers to test Go structs.
package structs

import (
	"reflect"
	"testing"

	"knative.dev/pkg/controller"
)

// EnsureNoNilField fails the test if the provided Impl's reconciler contains
// nil pointers or interfaces.
func EnsureNoNilField(t *testing.T, impl *controller.Impl) {
	t.Helper()

	recVal := reflect.ValueOf(impl.Reconciler).
		Elem().                    // injection/reconciler/sources/v1alpha1/<type>.reconcilerImpl
		FieldByName("reconciler"). // injection/reconciler/sources/v1alpha1/<type>.Interface
		Elem().                    // *pkg/reconciler/<type>.Reconciler (ptr)
		Elem()                     //  pkg/reconciler/<type>.Reconciler (val)

	for i := 0; i < recVal.NumField(); i++ {
		f := recVal.Field(i)
		switch f.Kind() {
		case reflect.Interface, reflect.Ptr, reflect.Func:
			if f.IsNil() {
				t.Errorf("struct field %q is nil", recVal.Type().Field(i).Name)
			}
		}
	}
}
