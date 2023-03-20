package findindexoffirststringoccurrence

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_strStr(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		idx := strStr("sadbutsad", "sad")
		require.Equal(t, 0, idx)
	})

	t.Run("2", func(t *testing.T) {
		idx := strStr("leetcode", "leeto")
		require.Equal(t, -1, idx)
	})

	t.Run("3", func(t *testing.T) {
		idx := strStr("hello", "ll")
		require.Equal(t, 2, idx)
	})

	t.Run("4", func(t *testing.T) {
		idx := strStr("a", "a")
		require.Equal(t, 0, idx)
	})

	t.Run("5", func(t *testing.T) {
		idx := strStr("abc", "c")
		require.Equal(t, 2, idx)
	})

	t.Run("6", func(t *testing.T) {
		idx := strStr("mississippi", "issi")
		require.Equal(t, 1, idx)
	})
}
