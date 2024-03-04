package backupstore

import (
	"encoding/json"
	"fmt"

	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/spf13/pflag"
)

type StoreType string

type DynamicOptions map[string]interface{}

func (o DynamicOptions) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(o)
	return data, err
}

type Options struct {
	Type     string         `json:"type" yaml:"type"`
	Provider DynamicOptions `json:"provider" yaml:"provider"`
}

func NewOptions() *Options {
	return &Options{
		Type:     "fs",
		Provider: map[string]interface{}{},
	}
}

func (o *Options) Validate() error {
	providers := sets.NewString()
	for key := range providerFactories {
		providers.Insert(key)
	}
	if !providers.Has(o.Type) {
		return fmt.Errorf("unsupport backend store type: %s", o.Type)
	}
	return nil
}

func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Type, "backupstore-type", o.Type, "backup store type: fs/s3")
}
