package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	iamv1 "github.com/kubeclipper/kubeclipper/pkg/scheme/iam/v1"
)

type UserLister interface {
	// List lists all Users in the indexer.
	List(selector labels.Selector) (ret []*iamv1.User, err error)
	// Get retrieves the User from the index for a given name.
	Get(name string) (*iamv1.User, error)
	UserListerExpansion
}

type userLister struct {
	indexer cache.Indexer
}

func NewUserLister(indexer cache.Indexer) UserLister {
	return &userLister{indexer: indexer}
}

func (c *userLister) List(selector labels.Selector) (ret []*iamv1.User, err error) {
	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*iamv1.User))
	})
	return ret, err
}

func (c *userLister) Get(name string) (*iamv1.User, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(iamv1.Resource("user"), name)
	}
	return obj.(*iamv1.User), nil
}
