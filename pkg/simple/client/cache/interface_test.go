package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func CacheCommonTest(t *testing.T, kv Interface) {
	const (
		key   = "foo"
		value = "bar"
	)

	exist, err := kv.Exist(key)
	if assert.NoError(t, err) {
		assert.False(t, exist)
	}

	err = kv.Set(key, value, time.Second*5)
	assert.NoError(t, err)

	exist, err = kv.Exist(key)
	if assert.NoError(t, err) {
		assert.True(t, exist)
	}

	getValue, err := kv.Get(key)
	if assert.NoError(t, err) {
		assert.Equal(t, value, getValue)
	}

	err = kv.Expire(key, time.Second*10)
	assert.NoError(t, err)

	getValue, err = kv.Get(key)
	if assert.NoError(t, err) {
		assert.Equal(t, value, getValue)
	}

	err = kv.Remove(key)
	assert.NoError(t, err)

	_, err = kv.Get(key)
	assert.True(t, IsNotExists(err))
}
