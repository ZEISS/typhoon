// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/flow/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// WorkerTransformationLister helps list WorkerTransformations.
// All objects returned here must be treated as read-only.
type WorkerTransformationLister interface {
	// List lists all WorkerTransformations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WorkerTransformation, err error)
	// WorkerTransformations returns an object that can list and get WorkerTransformations.
	WorkerTransformations(namespace string) WorkerTransformationNamespaceLister
	WorkerTransformationListerExpansion
}

// workerTransformationLister implements the WorkerTransformationLister interface.
type workerTransformationLister struct {
	listers.ResourceIndexer[*v1alpha1.WorkerTransformation]
}

// NewWorkerTransformationLister returns a new WorkerTransformationLister.
func NewWorkerTransformationLister(indexer cache.Indexer) WorkerTransformationLister {
	return &workerTransformationLister{listers.New[*v1alpha1.WorkerTransformation](indexer, v1alpha1.Resource("workertransformation"))}
}

// WorkerTransformations returns an object that can list and get WorkerTransformations.
func (s *workerTransformationLister) WorkerTransformations(namespace string) WorkerTransformationNamespaceLister {
	return workerTransformationNamespaceLister{listers.NewNamespaced[*v1alpha1.WorkerTransformation](s.ResourceIndexer, namespace)}
}

// WorkerTransformationNamespaceLister helps list and get WorkerTransformations.
// All objects returned here must be treated as read-only.
type WorkerTransformationNamespaceLister interface {
	// List lists all WorkerTransformations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WorkerTransformation, err error)
	// Get retrieves the WorkerTransformation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WorkerTransformation, error)
	WorkerTransformationNamespaceListerExpansion
}

// workerTransformationNamespaceLister implements the WorkerTransformationNamespaceLister
// interface.
type workerTransformationNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.WorkerTransformation]
}
