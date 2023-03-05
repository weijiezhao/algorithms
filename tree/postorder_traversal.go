package tree

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	postorder(root, &res)
	return res
}

func postorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	postorder(root.Left, nums)
	postorder(root.Right, nums)
	*nums = append(*nums, root.Val)
}
