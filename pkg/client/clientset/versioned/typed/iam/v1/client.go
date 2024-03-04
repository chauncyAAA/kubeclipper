package v1

import (
	"k8s.io/client-go/rest"

	iamv1 "github.com/kubeclipper/kubeclipper/pkg/scheme/iam/v1"

	"github.com/kubeclipper/kubeclipper/pkg/client/clientrest"

	"github.com/kubeclipper/kubeclipper/pkg/client/clientset/versioned/scheme"
)

var _ IamV1Interface = (*IamV1Client)(nil)

type IamV1Interface interface {
	RESTClient() rest.Interface
	TokensGetter
}

type IamV1Client struct {
	restClient rest.Interface
}

func New(c rest.Interface) IamV1Interface {
	return &IamV1Client{restClient: c}
}

func (c *IamV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

func (c *IamV1Client) Tokens() TokenInterface {
	return newTokens(c)
}

func NewForConfig(c *rest.Config) (*IamV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := clientrest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &IamV1Client{client}, nil
}

func setConfigDefaults(config *rest.Config) error {
	gv := iamv1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/api"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	// TODO: add custom user agent
	//if config.UserAgent == "" {
	//	config.UserAgent = rest.DefaultKubernetesUserAgent()
	//}

	return nil
}
