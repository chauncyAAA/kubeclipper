package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/iam/v1"
)

type TokenLister interface {
	// List lists all Regions in the indexer.
	List(selector labels.Selector) (ret []*v1.Token, err error)
	// Get retrieves the Region from the index for a given name.
	Get(name string) (*v1.Token, error)
	TokenListerExpansion
}

type tokenLister struct {
	indexer cache.Indexer
}

func NewTokenLister(indexer cache.Indexer) TokenLister {
	return &tokenLister{indexer: indexer}
}

func (c *tokenLister) List(selector labels.Selector) (ret []*v1.Token, err error) {
	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Token))
	})
	return ret, err
}

func (c *tokenLister) Get(name string) (*v1.Token, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("token"), name)
	}
	return obj.(*v1.Token), nil
}
