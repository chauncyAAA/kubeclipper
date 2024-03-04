package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

var _ ClusterLister = (*clusterLister)(nil)

type ClusterLister interface {
	// List lists all Users in the indexer.
	List(selector labels.Selector) (ret []*v1.Cluster, err error)
	// Get retrieves the User from the index for a given name.
	Get(name string) (*v1.Cluster, error)
	ClusterListerExpansion
}

type clusterLister struct {
	indexer cache.Indexer
}

func NewClusterLister(indexer cache.Indexer) ClusterLister {
	return &clusterLister{indexer: indexer}
}

func (c *clusterLister) List(selector labels.Selector) (ret []*v1.Cluster, err error) {
	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Cluster))
	})
	return ret, err
}

func (c *clusterLister) Get(name string) (*v1.Cluster, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("cluster"), name)
	}
	return obj.(*v1.Cluster), nil
}
