package misc

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestIntPow(t *testing.T) {
	n := IntPow(2, 10)
	require.True(t, n == 1024)
	n = IntPow(23, 10)
	require.True(t, n == 41426511213649)
}

func TestGenRandBytes(t *testing.T) {
	rb1, err := GenRandBytes(1024)
	require.NoError(t, err)
	require.True(t, len(rb1) == 1024)
	require.False(t, bytes.Equal(rb1, []byte{}))

	rb2, err := GenRandBytes(1024)
	require.NoError(t, err)
	require.True(t, len(rb2) == 1024)
	require.False(t, bytes.Equal(rb2, []byte{}))

	require.False(t, bytes.Equal(rb1, rb2))
}

func TestGenRandBlocks(t *testing.T) {

	// BLOCK_SIZE := 1024 // 1 Kib
	BLOCK_SIZE := 1024 * 1024 // 1 Mib
	BLOCK_COUNT := 512

	start := time.Now()

	blocks, err := GenRandBlocks(BLOCK_SIZE, BLOCK_COUNT)
	require.NoError(t, err)

	require.True(t, len(blocks) == BLOCK_COUNT)
	require.True(t, len(blocks[0]) == BLOCK_SIZE)
	for i := 0; i < BLOCK_COUNT-1; i++ {
		require.False(t, bytes.Equal(blocks[i], blocks[i+1]))
	}

	elapsed := time.Since(start)
	t.Log("Time taken: ", elapsed, "to generate", BLOCK_SIZE*BLOCK_COUNT/1024/1024, "MiB")
}
