/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	head2 := reverseList(slow.Next)
	slow.Next = nil 

	l1, l2 := head, head2
	for l2 != nil {
		next1, next2 := l1.Next, l2.Next
		l1.Next = l2
		l2.Next = next1

		l1 = next1
		l2 = next2
	}
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	curr:= head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
