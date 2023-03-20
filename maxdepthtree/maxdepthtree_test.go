package maxdepthtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_merge(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		root := TreeNode{
			Val: -8,
			Left: &TreeNode{
				Val: -6,
				Left: &TreeNode{
					Val: 6,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			Right: &TreeNode{
				Val: 7,
			},
		}
		max := maxDepth(&root)
		require.Equal(t, 4, max)
	})
}
