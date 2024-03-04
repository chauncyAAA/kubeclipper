package core

import (
	"k8s.io/client-go/informers/internalinterfaces"

	"github.com/kubeclipper/kubeclipper/pkg/client/internal"

	v1 "github.com/kubeclipper/kubeclipper/pkg/client/informers/core/v1"
)

type Interface interface {
	V1() v1.Interface
}

type group struct {
	factor           internal.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func New(f internal.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &group{factor: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

func (g *group) V1() v1.Interface {
	return v1.New(g.factor, g.namespace, g.tweakListOptions)
}
