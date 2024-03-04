package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

var _ OperationLister = (*operationLister)(nil)

type OperationLister interface {
	// List lists all Users in the indexer.
	List(selector labels.Selector) (ret []*v1.Operation, err error)
	// Get retrieves the User from the index for a given name.
	Get(name string) (*v1.Operation, error)
	OperationListerExpansion
}

type operationLister struct {
	indexer cache.Indexer
}

func NewOperationLister(indexer cache.Indexer) OperationLister {
	return &operationLister{indexer: indexer}
}

func (c *operationLister) List(selector labels.Selector) (ret []*v1.Operation, err error) {
	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Operation))
	})
	return ret, err
}

func (c *operationLister) Get(name string) (*v1.Operation, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("operation"), name)
	}
	return obj.(*v1.Operation), nil
}
