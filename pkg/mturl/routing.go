package mturl

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// URLPath returns a URL path to route requests for the given object.
func URLPath(o metav1.Object) string {
	return "/" + o.GetNamespace() + "/" + o.GetName()
}
