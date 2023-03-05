package listnode

// BM10 两个链表的第一个公共结点

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil || pHead2 == nil {
		return nil
	}

	curNode1, curNode2 := pHead1, pHead2
	for curNode1 != curNode2 {
		curNode1 = curNode1.Next
		curNode2 = curNode2.Next

		if curNode1 == nil {
			curNode1 = pHead2
		}

		if curNode2 == nil {
			curNode2 = pHead1
		}
	}
	return curNode1
}
