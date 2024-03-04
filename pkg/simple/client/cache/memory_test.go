package cache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMemoryKVStorage(t *testing.T) {
	kv, err := NewMemory()
	require.NoError(t, err)
	CacheCommonTest(t, kv)
}
