/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
  maxSum := root.Val
  var dfs func(node *TreeNode) int
  dfs = func(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := max(0, dfs(node.Left))
	right := max (0, dfs(node.Right))
	current := node.Val + left + right 
	if current > maxSum{
		maxSum = current
	}
	return node.Val + max(left,right)
  }
  dfs(root)
  return maxSum
}


