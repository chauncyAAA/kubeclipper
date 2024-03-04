package clientset

import (
	"fmt"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"

	corev1 "github.com/kubeclipper/kubeclipper/pkg/client/clientset/versioned/typed/core/v1"
	iamv1 "github.com/kubeclipper/kubeclipper/pkg/client/clientset/versioned/typed/iam/v1"
)

type Interface interface {
	CoreV1() corev1.CoreV1Interface
	IamV1() iamv1.IamV1Interface
}

type Clientset struct {
	corev1 *corev1.CoreV1Client
	iamv1  *iamv1.IamV1Client
}

func (c *Clientset) CoreV1() corev1.CoreV1Interface {
	return c.corev1
}

func (c *Clientset) IamV1() iamv1.IamV1Interface {
	return c.iamv1
}

func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.corev1, err = corev1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.iamv1, err = iamv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}
