package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

var _ RegionLister = (*regionLister)(nil)

type RegionLister interface {
	// List lists all Regions in the indexer.
	List(selector labels.Selector) (ret []*v1.Region, err error)
	// Get retrieves the Region from the index for a given name.
	Get(name string) (*v1.Region, error)
	RegionListerExpansion
}

type regionLister struct {
	indexer cache.Indexer
}

func NewRegionLister(indexer cache.Indexer) RegionLister {
	return &regionLister{indexer: indexer}
}

func (c *regionLister) List(selector labels.Selector) (ret []*v1.Region, err error) {
	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Region))
	})
	return ret, err
}

func (c *regionLister) Get(name string) (*v1.Region, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("region"), name)
	}
	return obj.(*v1.Region), nil
}
