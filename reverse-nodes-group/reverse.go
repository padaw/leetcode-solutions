type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(start *ListNode, end *ListNode, k int) (*ListNode, *ListNode) {
	prev := end.Next
	curr := start

	for i := 0; i < k; i++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		if i < k-1 {
			curr = next
		}
	}

	return curr, start
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	base := &ListNode{Next: head}
	pre := base
	curr := head
	for i := 1; curr != nil; i++ {
		if i%k == 0 {
			start, end := reverse(pre.Next, curr, k)
			pre.Next = start
			pre = end
			curr = end.Next
		} else {
			curr = curr.Next
		}
	}
	return base.Next
}
