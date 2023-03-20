package finddeuplicates

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_merge(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		input := []int{1, 2, 3, 1}
		res := containsDuplicate(input)
		require.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		res := containsDuplicate(input)
		require.Equal(t, false, res)
	})
}
