package removeduplicates

import "testing"

func Test_RemoveDuplicates(t *testing.T) {
	t.Run("simple array", func(t *testing.T) {
		RemoveDuplicates([]int{1, 1, 2})
	})
}
