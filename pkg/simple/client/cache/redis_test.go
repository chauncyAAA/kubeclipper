package cache

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func getRedisOptionFromEnv(t *testing.T) *RedisOptions {
	var (
		err       error
		dbIndex   int
		Addresses []string
	)
	if s := os.Getenv("REDIS_DB"); s != "" {
		dbIndex, err = strconv.Atoi(s)
		if err != nil {
			t.Errorf("par")
			t.FailNow()
		}
	}

	if addrs := os.Getenv("REDIS_ADDRESSES"); addrs != "" {
		Addresses = strings.Split(addrs, ",")
	}

	opts := &RedisOptions{
		Schema:           os.Getenv("REDIS_SCHEMA"),
		Addrs:            Addresses,
		Username:         os.Getenv("REDIS_USERNAME"),
		Password:         os.Getenv("REDIS_PASSWORD"),
		DB:               dbIndex,
		MasterName:       os.Getenv("REDIS_MASTER_NAME"),
		SentinelUsername: os.Getenv("REDIS_SENTINEL_USERNAME"),
		SentinelPassword: os.Getenv("REDIS_SENTINEL_USERNAME"),
		Prefix:           fmt.Sprintf("unittest-%s", strconv.FormatUint(uint64(time.Now().Unix()), 36)),
	}

	if opts.Schema == "" {
		t.Skipf("not get REDIS_SCHEMA from environment variables")
	}
	if len(opts.Addrs) == 0 {
		t.Skipf("not get REDIS_ADDRESSES from environment variables")
	}
	return opts
}

func TestRedisKVStorage(t *testing.T) {
	kv, err := NewRedis(getRedisOptionFromEnv(t))
	require.NoError(t, err)
	CacheCommonTest(t, kv)
}
