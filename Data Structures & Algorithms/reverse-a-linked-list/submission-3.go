/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	// Now:
	//  	0 -> 	1 -> 	2 ->
	//
	// Expected:
	//		<- 0	<- 1	<- 2
	//
	//
	// Simulation:
	//
	//	 	0 -> 	1 -> 	2 ->
	//	t	h	
	//	
	// Store next head to the temporary place (next)
	//
	//	 	0 -> 	1 -> 	2 ->
	//	t	h		next
	// 
	// Reverse the current head.next with current tail. Since the initial tail is nil,
	// now current head.next is nil. This is fullfil the expectation.
	//
	//	 	<- 0	1 -> 	2 ->
	//	t	h		next
	// 
	// Traverse to the next head. 
	// Traverse tail to the current head.
	//
	//	 	<- 0	1 -> 	2 ->
	//		h		next
	//		t
	//
	// Take the temporary place (next) as the new head.
	//
	//	 	<- 0	1 -> 	2 ->
	//		h		next
	//		t		h
	//
	// You must traverse tail first in this scenario because head is tail need to
	// traverse to the old head, not the new head
	//
	var next, tail *ListNode
	for head != nil {
		next = head.Next
		head.Next = tail
		tail = head
		head = next
	}

	return tail
}
