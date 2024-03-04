package cache

import (
	"fmt"

	"github.com/spf13/pflag"
)

const (
	ProviderEtcd   = "etcd"
	ProviderRedis  = "redis"
	ProviderMemory = "memory"
)

type Options struct {
	CacheProvider string        `json:"cacheProvider" yaml:"cacheProvider"`
	RedisOptions  *RedisOptions `json:"redis,omitempty" yaml:"redis,omitempty" mapstructure:"redis"`
}

func NewEtcdOptions() *Options {
	return &Options{
		CacheProvider: ProviderEtcd,
		RedisOptions:  NewRedisOptions(),
	}
}

func (s *Options) Validate() []error {
	if s == nil {
		return nil
	}
	var errs []error
	if s.CacheProvider != ProviderEtcd &&
		s.CacheProvider != ProviderRedis &&
		s.CacheProvider != ProviderMemory {
		errs = append(errs, fmt.Errorf("not support cache provider:%s", s.CacheProvider))
	}
	errs = append(errs, s.RedisOptions.Validate()...)

	return errs
}

func (s *Options) AddFlags(fs *pflag.FlagSet) {
	if s == nil {
		return
	}
	fs.StringVar(&s.CacheProvider, "cache-provider", s.CacheProvider, "Cache underlying provider, must be one of 'etcd' or 'redis'.")
	s.RedisOptions.AddFlags(fs)
}
