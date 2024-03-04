package oplog

import (
	"errors"
	"path/filepath"

	"github.com/spf13/pflag"
)

const (
	DefaultDir       = "/var/log/kubeclipper-agent/operations"
	MaximumThreshold = 1048576 // 1MB
	DefaultThreshold = 1048576 // 1MB
)

type Options struct {
	Dir             string `json:"dir" yaml:"dir"`
	SingleThreshold int64  `json:"singleThreshold" yaml:"singleThreshold"`
}

func NewOptions() *Options {
	return &Options{
		Dir:             DefaultDir,
		SingleThreshold: DefaultThreshold,
	}
}

func (s *Options) Validate() (errs []error) {
	if s == nil {
		return nil
	}
	if s.Dir == "" {
		return append(errs, errors.New("the dir for storing operation logs is not specified"))
	}
	if s.Dir == "/" {
		return append(errs, errors.New("the dir for storing operation logs cannot be root path"))
	}
	if !filepath.IsAbs(s.Dir) {
		return append(errs, errors.New("the dir for storing operation logs must be absolute"))
	}
	if s.SingleThreshold <= 0 {
		return append(errs, errors.New("the threshold must be greater than 0"))
	}
	if s.SingleThreshold > MaximumThreshold {
		return append(errs, errors.New("the threshold exceeded the limit, the maximum threshold is 1MB"))
	}
	return
}

func (s *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.Dir, "oplog-dir", s.Dir, "directory of op log file")
	fs.StringVar(&s.Dir, "oplog-threshold", s.Dir, "maximum value of log data transfer")
}
