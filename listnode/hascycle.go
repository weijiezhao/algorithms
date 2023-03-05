package listnode

//BM6 判断链表中是否有环
func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for head != nil && head.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
