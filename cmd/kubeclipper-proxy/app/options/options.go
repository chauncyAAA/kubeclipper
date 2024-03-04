package options

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/kubeclipper/kubeclipper/pkg/proxy/config"
)

type ProxyOptions struct {
	*config.Config
}

func NewProxyOptions() *ProxyOptions {
	return &ProxyOptions{}
}

func (s *ProxyOptions) CompletionOptions() error {
	conf, err := config.TryLoadFromDisk()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// always use config file
			return fmt.Errorf("config file not found, now only support config file to init server")
		}
		return fmt.Errorf("error parsing configuration file %s", err)
	}
	s.Config = conf
	return nil
}

func (s *ProxyOptions) Validate() []error {
	var errors []error
	if s.DefaultMQPort == 0 {
		errors = append(errors, fmt.Errorf("defaultMQPort can't be empty"))
	}
	if s.DefaultStaticServerPort == 0 {
		errors = append(errors, fmt.Errorf("defaultStaticServerPort can't be empty"))
	}

	if s.DefaultMQPort != 0 && s.DefaultMQPort == s.DefaultStaticServerPort {
		errors = append(errors, fmt.Errorf("defaultMQPort(%v) can't equal defaultStaticServerPort(%v)", s.DefaultMQPort, s.DefaultStaticServerPort))
	}
	return errors
}
