/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// Linked list
	// 0 -> 1 -> 2 -> 3
	// 0 <- 1 <- 2 <- 3
	// 
	// To reverse, get with two pointers. Initial node and the next node
	//
	// h is head
	// t is tail
	// tmp is temporary head
	//
	// 1st loop:
	// nil  0 -> 1 -> 2 -> 3
	// t 	h	 tmp
	// - t = nil		nil
	// - initial h		0
	// - tmp = h.next	1
	// - h.next = t		nil
	// - continue
	// 
	//
	// 2nd loop:
	// nil <- 0 ... 1 -> 2 -> 3
	// 		  t     h	 tmp
	// - t = h			0
	// - h = tmp		1
	// - tmp = h.next	2
	// - h.next = t		0
	// - continue
	//
	
	var tail, tmp *ListNode
	for {
		tmp = head.Next
		head.Next = tail

		if tmp == nil {
			break
		}

		tail = head
		head = tmp
	}

	return head
}
