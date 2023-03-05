package listnode

// BM4 合并两个排序的链表

func Merge1(pHead1 *ListNode, pHead2 *ListNode) *ListNode {

	var newHead = new(ListNode) // 如果用 var newHead *ListNode 则 panic！why?
	cur := newHead
	for pHead1 != nil || pHead2 != nil {
		if pHead1 == nil {
			cur.Next = pHead2
			break
		}

		if pHead2 == nil {
			cur.Next = pHead1
			break
		}

		if pHead1.Val < pHead2.Val {
			cur.Next = pHead1
			cur = cur.Next
			pHead1 = pHead1.Next
			continue
		}
		cur.Next = pHead2
		cur = cur.Next
		pHead2 = pHead2.Next
	}

	return newHead.Next
}

func Merge2(pHead1 *ListNode, pHead2 *ListNode) *ListNode {

	if pHead1 == nil {
		return pHead2
	}

	if pHead2 == nil {
		return pHead1
	}

	if pHead1.Val < pHead2.Val {
		pHead1.Next = Merge2(pHead1.Next, pHead2)
		return pHead1
	}

	pHead2.Next = Merge2(pHead1, pHead2.Next)
	return pHead2
}
