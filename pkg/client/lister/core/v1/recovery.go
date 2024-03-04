package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

var _ RecoveryLister = (*recoveryLister)(nil)

type RecoveryLister interface {
	// List lists all Recoveries in the indexer.
	List(selector labels.Selector) (ret []*v1.Recovery, err error)
	// Get retrieves the Recovery from the index for a given name.
	Get(name string) (*v1.Recovery, error)
	RecoveryListerExpansion
}

type recoveryLister struct {
	indexer cache.Indexer
}

func NewRecoveryLister(indexer cache.Indexer) RecoveryLister {
	return &recoveryLister{indexer: indexer}
}

func (l *recoveryLister) List(selector labels.Selector) (ret []*v1.Recovery, err error) {
	err = cache.ListAll(l.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Recovery))
	})
	return ret, err
}

func (l *recoveryLister) Get(name string) (*v1.Recovery, error) {
	obj, exists, err := l.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("recovery"), name)
	}
	return obj.(*v1.Recovery), nil
}
