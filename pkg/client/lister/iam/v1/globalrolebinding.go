package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	iamv1 "github.com/kubeclipper/kubeclipper/pkg/scheme/iam/v1"
)

type GlobalRoleBindingLister interface {
	// List lists all Users in the indexer.
	List(selector labels.Selector) (ret []*iamv1.GlobalRoleBinding, err error)
	// Get retrieves the User from the index for a given name.
	Get(name string) (*iamv1.GlobalRoleBinding, error)
	GlobalRoleBindingListerExpansion
}

type globalRoleBindingLister struct {
	indexer cache.Indexer
}

func NewGlobalRoleBindingLister(indexer cache.Indexer) GlobalRoleBindingLister {
	return &globalRoleBindingLister{indexer: indexer}
}

func (c *globalRoleBindingLister) List(selector labels.Selector) (ret []*iamv1.GlobalRoleBinding, err error) {
	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*iamv1.GlobalRoleBinding))
	})
	return ret, err
}

func (c *globalRoleBindingLister) Get(name string) (*iamv1.GlobalRoleBinding, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(iamv1.Resource("globalrolebinding"), name)
	}
	return obj.(*iamv1.GlobalRoleBinding), nil
}
