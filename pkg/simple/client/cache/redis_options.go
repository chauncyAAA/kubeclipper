package cache

import (
	"github.com/spf13/pflag"
)

const (
	Redis         = "redis"
	RedisSentinel = "redis-sentinel"
	RedisCluster  = "redis-cluster"
)

type RedisOptions struct {
	// redis schema. one of redis redis-sentinel cluster
	Schema string `json:"schema"`

	Addrs    []string `json:"addrs" yaml:"addrs"`
	Username string   `json:"username" yaml:"username"`
	Password string   `json:"password" yaml:"password"`
	DB       int      `json:"db" yaml:"db"`

	MasterName       string `json:"masterName" yaml:"masterName"`
	SentinelUsername string `json:"sentinelUsername" yaml:"sentinelUsername"`
	SentinelPassword string `json:"sentinelPassword" yaml:"sentinelPassword"`

	Prefix string `json:"prefix" yaml:"prefix"`
}

func NewRedisOptions() *RedisOptions {
	return &RedisOptions{
		Schema: "redis",
		Prefix: "kubeclipper:cache",
	}
}

func (s *RedisOptions) Validate() []error {
	if s == nil {
		return nil
	}
	// TODO validate options...

	return nil
}

// AddFlags specified FlagSet
func (s *RedisOptions) AddFlags(fs *pflag.FlagSet) {
	if s == nil {
		return
	}

	fs.StringVar(&s.Schema, "redis-schema", s.Schema,
		"Redis schema. Must be one of redis, redis sentinel and cluster.")
	fs.StringSliceVar(&s.Addrs, "redis-addrs", s.Addrs,
		"A seed list of host:port addresses of redis nodes.")
	fs.StringVar(&s.Username, "redis-username", s.Username,
		"Redis username")
	fs.StringVar(&s.Password, "redis-password", s.Password,
		"Redis password.")
	fs.IntVar(&s.DB, "redis-db", s.DB,
		"Database to be selected after connecting to the redis.")
	fs.StringVar(&s.MasterName, "redis-master-name", s.MasterName,
		"Redis sentinel masterName")
	fs.StringVar(&s.SentinelUsername, "redis-sentinel-username", s.SentinelUsername,
		"Redis sentinel username.")
	fs.StringVar(&s.SentinelPassword, "redis-sentinel-password", s.SentinelPassword,
		"Redis sentinel password.")
	fs.StringVar(&s.Prefix, "redis-prefix", s.Prefix,
		"All redis keys will be prefixed.")

}
