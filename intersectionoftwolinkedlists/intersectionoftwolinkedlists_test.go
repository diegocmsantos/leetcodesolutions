package intersectionoftwolinkedlists

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		headA := buildNodes([]int{4, 1, 8, 4, 5})
		headB := buildNodes([]int{5, 6, 1, 8, 4, 5})
		res := getIntersectionNode(headA, headB)
		fmt.Println(res)
	})
}

func buildNodes(a []int) *ListNode {
	res := &ListNode{Val: a[0]}
	curr := res
	for _, n := range a[1:] {
		newNode := &ListNode{
			Val: n,
		}
		curr.Next = newNode
		curr = curr.Next
	}
	return res
}
