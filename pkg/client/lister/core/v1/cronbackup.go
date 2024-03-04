package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

var _ CronBackupLister = (*cronBackupLister)(nil)

type CronBackupLister interface {
	// List lists all CronBackups in the indexer.
	List(selector labels.Selector) (ret []*v1.CronBackup, err error)
	// Get retrieves the CronBackup from the index for a given name.
	Get(name string) (*v1.CronBackup, error)
	CronBackupListerExpansion
}

type cronBackupLister struct {
	indexer cache.Indexer
}

func NewCronBackupLister(indexer cache.Indexer) CronBackupLister {
	return &cronBackupLister{indexer: indexer}
}

func (l *cronBackupLister) List(selector labels.Selector) (ret []*v1.CronBackup, err error) {
	err = cache.ListAll(l.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CronBackup))
	})
	return ret, err
}

func (l *cronBackupLister) Get(name string) (*v1.CronBackup, error) {
	obj, exists, err := l.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("cronbackup"), name)
	}
	return obj.(*v1.CronBackup), nil
}
