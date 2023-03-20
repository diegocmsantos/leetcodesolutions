package intersectionoftwolinkedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for headA != nil || headB != nil {
		if headA != nil {
			if _, ok := m[headA]; ok {
				return headA
			}
			m[headA] = struct{}{}
			headA = headA.Next
		}
		if headB != nil {
			if _, ok := m[headB]; ok {
				return headB
			}
			m[headB] = struct{}{}
			headB = headB.Next
		}
	}
	return nil
}
