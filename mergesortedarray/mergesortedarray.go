package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m+n-1; j++ {
			if nums1[j] == 0 || (nums2[i] >= nums1[j] && nums2[i] < nums1[j+1]) {
				nums1 = insert(nums1, nums2[i], j)
				break
			}
		}
	}
}

func insert(a []int, c int, i int) []int {
	return append(a[:i], append([]int{c}, a[i:len(a)-1]...)...)
}
