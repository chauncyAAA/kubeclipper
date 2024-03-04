package staticserver

import (
	"fmt"
	"os"

	"github.com/kubeclipper/kubeclipper/pkg/utils/netutil"
)

// TODO: add static server auth

type Options struct {
	BindAddress   string `json:"bindAddress" yaml:"bindAddress"`
	InsecurePort  int    `json:"insecurePort" yaml:"insecurePort"`
	SecurePort    int    `json:"securePort" yaml:"securePort"`
	TLSCertFile   string `json:"tlsCertFile" yaml:"tlsCertFile"`
	TLSPrivateKey string `json:"tlsPrivateKey" yaml:"tlsPrivateKey"`
	Path          string `json:"path" yaml:"path"`
}

func NewOptions() *Options {
	s := Options{
		BindAddress:   "0.0.0.0",
		InsecurePort:  8090,
		SecurePort:    0,
		TLSCertFile:   "",
		TLSPrivateKey: "",
		Path:          "/opt/kubeclipper-server/resource",
	}
	return &s
}

func (s *Options) Validate() []error {
	var errs []error

	if s.SecurePort == 0 && s.InsecurePort == 0 {
		errs = append(errs, fmt.Errorf("insecure and secure port can not be disabled at the same time"))
	}

	if netutil.IsValidPort(s.SecurePort) {
		if s.TLSCertFile == "" {
			errs = append(errs, fmt.Errorf("tls cert file is empty while secure serving"))
		} else {
			if _, err := os.Stat(s.TLSCertFile); err != nil {
				errs = append(errs, err)
			}
		}

		if s.TLSPrivateKey == "" {
			errs = append(errs, fmt.Errorf("tls private key file is empty while secure serving"))
		} else {
			if _, err := os.Stat(s.TLSPrivateKey); err != nil {
				errs = append(errs, err)
			}
		}
	}

	if s.Path == "" {
		errs = append(errs, fmt.Errorf("static server resource path can not be empty"))
	}

	return errs
}

type AgentOptions struct {
	Address       string `json:"address" yaml:"address"`
	TLSCertFile   string `json:"tlsCertFile" yaml:"tlsCertFile"`
	TLSPrivateKey string `json:"tlsPrivateKey" yaml:"tlsPrivateKey"`
}

func NewAgentOptions() *AgentOptions {
	s := AgentOptions{
		Address:       "",
		TLSCertFile:   "",
		TLSPrivateKey: "",
	}
	return &s
}

func (s *AgentOptions) Validate() []error {
	return nil
}
