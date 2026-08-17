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

    list := make([]*ListNode, 0)

	node := head
	for {
		list = append(list, node)

		if node.Next != nil {
			node = node.Next
		} else if node.Next == nil {
			break
		}
	}

	// debug
	// for _, n := range list {
	// 	fmt.Println("node:", n)
	// }

	for i := len(list)-1; i >= 0; i-- {
		node = list[i]
		fmt.Println("head:", node)

		if i == 0 {
			node.Next = nil
		} else {
			node.Next = list[i-1]
		}
	}

	// fmt.Println("head:", list[len(list)-1])

	return list[len(list)-1]
}
