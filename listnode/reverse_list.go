package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	var newHead *ListNode
	for pHead != nil {
		pNext := pHead.Next
		pHead.Next = newHead
		newHead = pHead
		pHead = pNext
	}

	return newHead
}
