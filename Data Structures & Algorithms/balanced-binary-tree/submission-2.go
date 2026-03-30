/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    return dfs(root).balanced
}
type Result struct{
	balanced bool
	height int
}
func dfs(root *TreeNode) Result{
	if root == nil{
		return Result{true, 0}
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	balanced := left.balanced && right.balanced && abs(left.height - right.height) <= 1
	return Result{balanced, 1 + max(left.height, right.height)}
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}