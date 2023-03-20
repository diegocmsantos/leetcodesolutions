package mergesortedarray

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 0, 0, 0}
		m := 3
		nums2 := []int{2, 5, 6}
		n := 3
		merge(nums1, m, nums2, n)
		fmt.Println(nums1)
	})
}
