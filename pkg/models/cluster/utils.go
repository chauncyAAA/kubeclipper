package cluster

import (
	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

// Deprecated
type recordList []v1.Record

func (l recordList) Len() int      { return len(l) }
func (l recordList) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l recordList) Less(i, j int) bool {
	ti := l[i].CreateTime.Time
	tj := l[j].CreateTime.Time
	if ti.After(tj) {
		return true
	} else if ti.Before(tj) {
		return false
	} else {
		return l[i].RR > l[j].RR
	}
}
