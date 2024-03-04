package scheme

import (
	corev1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
	iamv1 "github.com/kubeclipper/kubeclipper/pkg/scheme/iam/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
)

var (
	// Scheme defines methods for serializing and deserializing API objects.
	Scheme = runtime.NewScheme()
	// Codecs provides methods for retrieving codecs and serializers for specific
	// versions and content types.
	Codecs = serializer.NewCodecFactory(Scheme)
	// Encoder provides json encoder
	Encoder = json.NewSerializerWithOptions(
		json.DefaultMetaFactory, Scheme, Scheme,
		json.SerializerOptions{Yaml: false, Pretty: false, Strict: false},
	)
)

func init() {
	_ = corev1.AddToScheme(Scheme)
	_ = iamv1.AddToScheme(Scheme)
}
