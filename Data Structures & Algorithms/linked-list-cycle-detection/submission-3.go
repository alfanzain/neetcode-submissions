/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	normalPointer, fastPointer := head, head

	if normalPointer == nil || normalPointer.Next == nil {
		return false
	}

	fastPointer = fastPointer.Next.Next
	normalPointer = normalPointer.Next

	for normalPointer != fastPointer {
		// fmt.Println("fast:", fastPointer.Val)
		// fmt.Println("normal:", normalPointer.Val)
		if fastPointer == nil || fastPointer.Next == nil {
			return false
		}

		fastPointer = fastPointer.Next.Next
		normalPointer = normalPointer.Next
	}

	return true
}
