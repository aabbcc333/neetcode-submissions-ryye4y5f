/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}

	res := abs(maxDepth(root.Left), maxDepth(root.Right))
	if res > 1{
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
	
}

func maxDepth(root *TreeNode) int{
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
func abs(a,b int) int{
	if a > b{
		return a-b
	}else{
		return b-a 
	}
}
