package iam

import iamv1 "github.com/kubeclipper/kubeclipper/pkg/scheme/iam/v1"

type userList []iamv1.User

func (l userList) Len() int      { return len(l) }
func (l userList) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l userList) Less(i, j int) bool {
	t1 := l[i].CreationTimestamp.Time
	t2 := l[j].CreationTimestamp.Time
	if t1.After(t2) {
		return true
	} else if t1.Before(t2) {
		return false
	} else {
		return l[i].Name > l[j].Name
	}
}

// DesensitizationUserPassword mark user's passwd empty.
func DesensitizationUserPassword(user *iamv1.User) {
	user.Spec.EncryptedPassword = ""
}
