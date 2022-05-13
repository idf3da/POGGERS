package storage

import (
	"bytes"
	"testing"

	"github.com/idf3da/POGGERS/internal/misc"
	"github.com/stretchr/testify/require"
)

func TestAdd1(t *testing.T) {
	var store Store
	store.Lib = make(map[[32]byte][]byte)
	store.T = "Local"

	sent := []byte("Whatever?")

	hash, err := store.Add(sent)
	require.NoError(t, err)

	got, err := store.Get(hash)
	require.NoError(t, err)
	require.True(t, bytes.Equal(got, sent))
}

func TestAdd3Different(t *testing.T) {
	var store Store
	store.Lib = make(map[[32]byte][]byte)
	store.T = "Local"

	sent1 := []byte("Whatever?")
	sent2 := []byte("However?")
	sent3 := []byte("Whenever?")

	hash1, err := store.Add(sent1)
	require.NoError(t, err)
	hash2, err := store.Add(sent2)
	require.NoError(t, err)
	hash3, err := store.Add(sent3)
	require.NoError(t, err)

	got1, err := store.Get(hash1)
	require.NoError(t, err)
	require.True(t, bytes.Equal(got1, sent1))
	got2, err := store.Get(hash2)
	require.NoError(t, err)
	require.True(t, bytes.Equal(got2, sent2))
	got3, err := store.Get(hash3)
	require.NoError(t, err)
	require.True(t, bytes.Equal(got3, sent3))
}

func TestAdd3Same(t *testing.T) {
	var store Store
	store.Lib = make(map[[32]byte][]byte)
	store.T = "Local"

	sent1 := []byte("SameSame?")
	sent2 := []byte("SameSame?")
	sent3 := []byte("SameSame?")

	hash1, err := store.Add(sent1)
	require.NoError(t, err)
	hash2, err := store.Add(sent2)
	require.NoError(t, err)
	hash3, err := store.Add(sent3)
	require.NoError(t, err)

	got1, err := store.Get(hash1)
	require.NoError(t, err)
	require.True(t, bytes.Equal(got1, sent1))
	got2, err := store.Get(hash2)
	require.NoError(t, err)
	require.True(t, bytes.Equal(got2, sent2))
	got3, err := store.Get(hash3)
	require.NoError(t, err)
	require.True(t, bytes.Equal(got3, sent3))
}

func TestAdd512Mib(t *testing.T) {
	BLOCK_SIZE := 1024 * 1024
	BLOCK_COUNT := 512

	var store Store
	store.Lib = make(map[[32]byte][]byte)
	store.T = "Local"

	data, _ := misc.GenRandBlocks(BLOCK_SIZE, BLOCK_COUNT)

	h := make([][32]byte, 0, BLOCK_COUNT)

	for _, v := range data {
		hash, err := store.Add(v)
		require.NoError(t, err)
		h = append(h, hash)
	}

	for i := 0; i < len(h); i++ {
		el, ok := store.Lib[h[i]]
		require.True(t, ok)
		require.True(t, bytes.Equal(el, data[i]))
	}
}

// TODO: Use random data henerator, test hash collision
