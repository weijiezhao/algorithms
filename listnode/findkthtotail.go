package listnode

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	slow, fast := pHead, pHead
	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
