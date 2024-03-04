package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

var _ DomainLister = (*domainLister)(nil)

type DomainLister interface {
	// List lists all Users in the indexer.
	List(selector labels.Selector) (ret []*v1.Domain, err error)
	// Get retrieves the User from the index for a given name.
	Get(name string) (*v1.Domain, error)
	DomainListerExpansion
}

type domainLister struct {
	indexer cache.Indexer
}

func NewDomainLister(indexer cache.Indexer) DomainLister {
	return &domainLister{indexer: indexer}
}

func (c *domainLister) List(selector labels.Selector) (ret []*v1.Domain, err error) {
	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Domain))
	})
	return ret, err
}

func (c *domainLister) Get(name string) (*v1.Domain, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("domain"), name)
	}
	return obj.(*v1.Domain), nil
}
