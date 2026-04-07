/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }

    // Map original -> copy
    nodeMap := make(map[*Node]*Node)

    // 1st pass: create all nodes
    curr := head
    for curr != nil {
        nodeMap[curr] = &Node{Val: curr.Val}
        curr = curr.Next
    }

    // 2nd pass: assign next and random
    curr = head
    for curr != nil {
        copyNode := nodeMap[curr]

        if curr.Next != nil {
            copyNode.Next = nodeMap[curr.Next]
        }

        if curr.Random != nil {
            copyNode.Random = nodeMap[curr.Random]
        }

        curr = curr.Next
    }

    return nodeMap[head]
}