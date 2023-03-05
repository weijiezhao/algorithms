package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	preorder(root, &res)
	return res
}

func preorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	*nums = append(*nums, root.Val)
	preorder(root.Left, nums)
	preorder(root.Right, nums)
}
