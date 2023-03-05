package tree

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	preorder(root, &res)
	return res
}

func inorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inorder(root, nums)
	*nums = append(*nums, root.Val)
	inorder(root.Right, nums)
}
