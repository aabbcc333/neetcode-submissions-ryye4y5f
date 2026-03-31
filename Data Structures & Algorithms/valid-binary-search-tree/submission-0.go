/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    var prev *TreeNode

    var inorder func(node *TreeNode) bool
    inorder = func(node *TreeNode) bool {
        if node == nil {
            return true
        }

        if !inorder(node.Left) {
            return false
        }

        if prev != nil && node.Val <= prev.Val {
            return false
        }
        prev = node

        return inorder(node.Right)
    }

    return inorder(root)
}

