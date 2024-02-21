

package googlecloudworkflowstarget

import (
	"testing"

	. "github.com/zeiss/typhoon/pkg/reconciler/testing"

	// Link fake informers accessed by our controller
	_ "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/targets/v1alpha1/googlecloudworkflowstarget/fake"
	_ "knative.dev/pkg/client/injection/ducks/duck/v1/addressable/fake"
	_ "knative.dev/pkg/client/injection/kube/informers/core/v1/serviceaccount/fake"
	_ "knative.dev/pkg/client/injection/kube/informers/rbac/v1/rolebinding/fake"
	_ "knative.dev/serving/pkg/client/injection/informers/serving/v1/service/fake"
)

func TestNewController(t *testing.T) {
	t.Run("No failure", func(t *testing.T) {
		TestControllerConstructor(t, NewController)
	})

	t.Run("Failure cases", func(t *testing.T) {
		TestControllerConstructorFailures(t, NewController)
	})github.com/zeiss/typhoon
}
