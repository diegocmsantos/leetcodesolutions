package addtwonumbers

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{}

	var num1 string
	for {
		if l1.Next != nil {
			num1 += strconv.Itoa(l1.Val)
			l1 = l1.Next
		} else {
			num1 += strconv.Itoa(l1.Val)
			break
		}
	}

	fmt.Println(num1)

	var num2 string
	for {
		if l2.Next != nil {
			num2 += strconv.Itoa(l2.Val)
			l2 = l2.Next
		} else {
			num2 += strconv.Itoa(l2.Val)
			break
		}
	}

	fmt.Println(num2)

	return &result
}
