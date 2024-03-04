package v1

import (
	coordinationv1 "k8s.io/api/coordination/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

var _ LeaseLister = (*leaseLister)(nil)

type LeaseLister interface {
	// List lists all Users in the indexer.
	List(selector labels.Selector) (ret []*coordinationv1.Lease, err error)
	// Get retrieves the User from the index for a given name.
	Get(name string) (*coordinationv1.Lease, error)
	Leases(namespace string) LeaseNamespaceLister
	LeaseListerExpansion
}

type leaseLister struct {
	indexer cache.Indexer
}

func NewLeaseLister(indexer cache.Indexer) LeaseLister {
	return &leaseLister{indexer: indexer}
}

func (c *leaseLister) List(selector labels.Selector) (ret []*coordinationv1.Lease, err error) {
	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*coordinationv1.Lease))
	})
	return ret, err
}

func (c *leaseLister) Get(name string) (*coordinationv1.Lease, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("lease"), name)
	}
	return obj.(*coordinationv1.Lease), nil
}

// Leases returns an object that can list and get Leases.
func (c *leaseLister) Leases(namespace string) LeaseNamespaceLister {
	return leaseNamespaceLister{indexer: c.indexer, namespace: namespace}
}

type LeaseNamespaceLister interface {
	// List lists all Users in the indexer.
	List(selector labels.Selector) (ret []*coordinationv1.Lease, err error)
	// Get retrieves the User from the index for a given name.
	Get(name string) (*coordinationv1.Lease, error)
}

// leaseNamespaceLister implements the LeaseNamespaceLister
// interface.
type leaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Leases in the indexer for a given namespace.
func (s leaseNamespaceLister) List(selector labels.Selector) (ret []*coordinationv1.Lease, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*coordinationv1.Lease))
	})
	return ret, err
}

// Get retrieves the Lease from the indexer for a given namespace and name.
func (s leaseNamespaceLister) Get(name string) (*coordinationv1.Lease, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("lease"), name)
	}
	return obj.(*coordinationv1.Lease), nil
}
