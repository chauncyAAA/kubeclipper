package mfa

import (
	"fmt"

	"github.com/spf13/pflag"

	"github.com/kubeclipper/kubeclipper/pkg/authentication/oauth"
)

func NewOptions() *Options {
	return &Options{
		Enabled:      false,
		MFAProviders: nil,
	}
}

type Options struct {
	Enabled      bool              `json:"enabled" yaml:"enabled"`
	MFAProviders []ProviderOptions `json:"mfaProviders" yaml:"mfaProviders"`
}

type ProviderOptions struct {
	Type    string               `json:"type" yaml:"type"`
	Options oauth.DynamicOptions `json:"options" yaml:"options"`
}

func (a *Options) Validate() []error {
	var errs []error
	if a.Enabled && len(a.MFAProviders) == 0 {
		errs = append(errs, fmt.Errorf("mfa is not configured"))
	}
	return errs
}

func (a *Options) AddFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&a.Enabled, "mfa-enabled", a.Enabled, "Enable multi-factor authentication.")
}
