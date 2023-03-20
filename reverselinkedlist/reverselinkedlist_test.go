package reverselinkedlist

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		}
		res := reverseList(head)
		fmt.Println(res)
	})
}
