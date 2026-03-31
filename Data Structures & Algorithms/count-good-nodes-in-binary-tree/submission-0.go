/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
   var dfs func(node *TreeNode, max int) int
   
   dfs = func(node *TreeNode, max int)int{
	 if node == nil {
		return 0
	 }
	 count := 0 
	 if node.Val >= max{
		count = 1
		max = node.Val
	 }
	 count += dfs(node.Left, max)
	 count += dfs(node.Right, max)
	 return count 
   }
   return dfs(root, root.Val)

}
