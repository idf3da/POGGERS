package misc

import (
	"math/rand"
)

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func GenRandBytes(n int) ([]byte, error) {
	token := make([]byte, n)
	_, err := rand.Read(token)
	return token, err
}

func GenRandBlocks(size, count int) ([][]byte, error) {
	blocks := make([][]byte, 0, count)

	for i := 0; i < count; i++ {
		block, err := GenRandBytes(size)
		if err != nil {
			return blocks, err
		}
		blocks = append(blocks, block)
	}

	return blocks, nil
}
