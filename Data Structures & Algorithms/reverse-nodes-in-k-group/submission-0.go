/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/* 
 find kth node 
 reverse the segment
 connect :
 prev group -> new head
 old head -> next group
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{Next: head}
    groupPrev := dummy

    for {
        kth := getKth(groupPrev, k)
        if kth == nil {
            break
        }

        groupNext := kth.Next

        // reverse group
        prev := groupNext
        curr := groupPrev.Next

        for curr != groupNext {
            temp := curr.Next
            curr.Next = prev
            prev = curr
            curr = temp
        }

        temp := groupPrev.Next
        groupPrev.Next = kth
        groupPrev = temp
    }

    return dummy.Next
}

func getKth(curr *ListNode, k int) *ListNode {
    for curr != nil && k > 0 {
        curr = curr.Next
        k--
    }
    return curr
}