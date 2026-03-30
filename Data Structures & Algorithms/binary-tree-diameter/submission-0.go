/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
   if root == nil {
    return 0 
   }

   diameter := maxDepth(root.Left)+ maxDepth(root.Right)
   diameter = max(diameter, max(diameterOfBinaryTree(root.Left),diameterOfBinaryTree(root.Right)))
   return diameter
    }

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}