package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	iamv1 "github.com/kubeclipper/kubeclipper/pkg/scheme/iam/v1"
)

type GlobalRoleLister interface {
	// List lists all Users in the indexer.
	List(selector labels.Selector) (ret []*iamv1.GlobalRole, err error)
	// Get retrieves the User from the index for a given name.
	Get(name string) (*iamv1.GlobalRole, error)
	GlobalRoleListerExpansion
}

type globalRoleLister struct {
	indexer cache.Indexer
}

func NewGlobalRoleLister(indexer cache.Indexer) GlobalRoleLister {
	return &globalRoleLister{indexer: indexer}
}

func (c *globalRoleLister) List(selector labels.Selector) (ret []*iamv1.GlobalRole, err error) {
	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*iamv1.GlobalRole))
	})
	return ret, err
}

func (c *globalRoleLister) Get(name string) (*iamv1.GlobalRole, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(iamv1.Resource("globalrole"), name)
	}
	return obj.(*iamv1.GlobalRole), nil
}
