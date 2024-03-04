package v1

import (
	"k8s.io/client-go/informers/internalinterfaces"

	"github.com/kubeclipper/kubeclipper/pkg/client/internal"
)

type Interface interface {
	Tokens() TokenInformer
}

type version struct {
	factory          internal.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internal.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

func (v *version) Tokens() TokenInformer {
	return &tokenInformer{
		factory:          v.factory,
		tweakListOptions: v.tweakListOptions,
	}
}
