/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    indexMap := make(map[int]int)
	for i, val := range inorder{
		indexMap[val] = i
	}
	var build func(preStart, inStart, inEnd int) *TreeNode
	build = func(preStart, inStart, inEnd int) *TreeNode{
		if inStart > inEnd{
			return nil
		}
		rootVal := preorder[preStart]
		root := &TreeNode{Val : rootVal}

		inIndex := indexMap[rootVal]
		leftSize := inIndex - inStart

		root.Left = build(preStart +1 , inStart, inIndex -1)
		root.Right = build(preStart +1 + leftSize, inIndex +1 ,inEnd)
		return root
	}
	return build(0,0, len(inorder)-1)
}
