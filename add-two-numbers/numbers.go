type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := &ListNode{Val: 0}
	res := tmp

	for l1 != nil || l2 != nil {
		if l1 != nil {
			tmp.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp.Val += l2.Val
			l2 = l2.Next
		}
		n := &ListNode{Val: 0}
		if tmp.Val > 9 {
			tmp.Val = tmp.Val % 10
			n.Val = 1
		}
		if l1 != nil || l2 != nil || n.Val > 0 {
			tmp.Next = n
			tmp = n
		}
	}

	return res
}
