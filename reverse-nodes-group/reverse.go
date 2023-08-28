type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(start *ListNode, end *ListNode, k int) (*ListNode, *ListNode) {
	arr := []*ListNode{}
	next := end.Next

	curr := start
	for i := 0; i < k; i++ {
		arr = append(arr, curr)
		curr = curr.Next
	}

	for i := k - 1; i >= 0; i-- {
		n := arr[i]
		if i > 0 {
			n.Next = arr[i-1]
		} else {
			n.Next = next
		}
	}

	return arr[k-1], arr[0]
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
