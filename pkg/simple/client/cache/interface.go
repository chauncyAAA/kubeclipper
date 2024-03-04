package cache

import (
	"errors"
	"fmt"
	"time"
)

const NoExpiration time.Duration = 0

var ErrNotExists = fmt.Errorf("key not exists")

type Interface interface {
	Set(key, value string, expire time.Duration) error
	Update(key, newValue string) error
	Get(key string) (value string, err error)
	Exist(key string) (bool, error)
	Remove(key string) error
	Expire(key string, expire time.Duration) error
}

func IsNotExists(e error) bool {
	return errors.Is(e, ErrNotExists)
}
