package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

var _ NodeLister = (*nodeLister)(nil)

type NodeLister interface {
	// List lists all Users in the indexer.
	List(selector labels.Selector) (ret []*v1.Node, err error)
	// Get retrieves the User from the index for a given name.
	Get(name string) (*v1.Node, error)
	NodeListerExpansion
}

type nodeLister struct {
	indexer cache.Indexer
}

func NewNodeLister(indexer cache.Indexer) NodeLister {
	return &nodeLister{indexer: indexer}
}

func (c *nodeLister) List(selector labels.Selector) (ret []*v1.Node, err error) {
	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Node))
	})
	return ret, err
}

func (c *nodeLister) Get(name string) (*v1.Node, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("node"), name)
	}
	return obj.(*v1.Node), nil
}
