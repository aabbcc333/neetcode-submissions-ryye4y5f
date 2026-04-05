/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 1. Find middle (slow/fast)
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. Reverse second half
	var prev *ListNode
	curr := slow.Next
	slow.Next = nil // split list

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// 3. Merge two halves
	first, second := head, prev

	for second != nil {
		tmp1 := first.Next
		tmp2 := second.Next

		first.Next = second
		second.Next = tmp1

		first = tmp1
		second = tmp2
	}
}