package scheme

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"

	corev1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
	iamv1 "github.com/kubeclipper/kubeclipper/pkg/scheme/iam/v1"
)

var (
	// Scheme defines methods for serializing and deserializing API objects.
	Scheme             = runtime.NewScheme()
	Codecs             = serializer.NewCodecFactory(Scheme)
	ParameterCodec     = runtime.NewParameterCodec(Scheme)
	localSchemeBuilder = runtime.SchemeBuilder{
		corev1.AddToScheme,
		iamv1.AddToScheme,
	}
	AddToScheme = localSchemeBuilder.AddToScheme
)

func init() {
	v1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	if err := AddToScheme(Scheme); err != nil {
		panic(err)
	}
}
