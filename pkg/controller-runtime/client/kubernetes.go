package client

import (
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/kubeclipper/kubeclipper/pkg/logger"
)

type Client interface {
	Kubernetes() kubernetes.Interface
	Config() *rest.Config
}

type kubernetesClient struct {
	// kubernetes client interface
	k8s    kubernetes.Interface
	config *rest.Config
}

// NewKubernetesClient creates a KubernetesClient
func NewKubernetesClient(config *rest.Config, client *kubernetes.Clientset) Client {
	return &kubernetesClient{
		k8s:    client,
		config: config,
	}
}

func (k *kubernetesClient) Kubernetes() kubernetes.Interface {
	return k.k8s
}

func (k *kubernetesClient) Config() *rest.Config {
	return k.config
}

func FromKubeConfig(kubeConfig []byte) (*rest.Config, *kubernetes.Clientset, error) {
	clientConfig, err := clientcmd.NewClientConfigFromBytes(kubeConfig)
	if err != nil {
		logger.Error("create cluster client config failed", zap.Error(err))
		return nil, nil, err
	}
	clientcfg, err := clientConfig.ClientConfig()
	if err != nil {
		logger.Error("get cluster kubeconfig client failed", zap.Error(err))
		return nil, nil, err
	}
	clientset, err := kubernetes.NewForConfig(clientcfg)
	if err != nil {
		logger.Error("create cluster clientset failed", zap.Error(err))
		return nil, nil, err
	}
	return clientcfg, clientset, nil
}
