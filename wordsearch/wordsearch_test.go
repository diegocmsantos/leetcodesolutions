package wordsearch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_merge(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		board := [][]byte{
			{byte('A'), byte('B'), byte('C'), byte('E')},
			{byte('S'), byte('F'), byte('C'), byte('S')},
			{byte('A'), byte('D'), byte('E'), byte('E')},
		}
		word := "ABCCED"
		exists := exist(board, word)
		require.Equal(t, true, exists)
	})
}
