package listnode

import "sort"

func sortInList(head *ListNode) *ListNode {
	curNode, newHead := head, head
	nums := make([]int, 0)
	for curNode != nil {
		nums = append(nums, curNode.Val)
		curNode = curNode.Next
	}

	sort.Ints(nums)

	for _, v := range nums {
		head.Val = v
		head = head.Next
	}
	return newHead
}
