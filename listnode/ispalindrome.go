package listnode

func isPalindrome(head *ListNode) bool {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[left] != nums[right] {
			return false
		}
	}
	return true
}
