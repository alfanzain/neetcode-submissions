// Singly Linked List Node
type ListNode struct {
	val  int
	next *ListNode
}

// NewListNode creates a new list node
func NewListNode(val int, nextNode *ListNode) *ListNode {
	return &ListNode{
		val:  val,
		next: nextNode,
	}
}

// Implementation for Singly Linked List
type LinkedList struct {
	head *ListNode
	tail *ListNode
}

func NewLinkedList() *LinkedList {
	// Init the list with a 'dummy' node which makes
	// removing a node from the beginning of list easier.
	head := NewListNode(-1, nil)
	return &LinkedList{
		head: head,
		tail: head,
	}
}

func (ll *LinkedList) Get(index int) int {
	curr := ll.head.next
	i := 0
	for curr != nil {
		if i == index {
			return curr.val
		}
		i++
		curr = curr.next
	}
	return -1 // Index out of bounds or list is empty
}

func (ll *LinkedList) InsertHead(val int) {
	newNode := NewListNode(val, ll.head.next)
	ll.head.next = newNode
	if newNode.next == nil { // If list was empty before insertion
		ll.tail = newNode
	}
}

func (ll *LinkedList) InsertTail(val int) {
	ll.tail.next = NewListNode(val, nil)
	ll.tail = ll.tail.next
}

func (ll *LinkedList) Remove(index int) bool {
	i := 0
	curr := ll.head
	for i < index && curr != nil {
		i++
		curr = curr.next
	}

	// Remove the node ahead of curr
	if curr != nil && curr.next != nil {
		if curr.next == ll.tail {
			ll.tail = curr
		}
		curr.next = curr.next.next
		return true
	}
	return false
}

func (ll *LinkedList) GetValues() []int {
	curr := ll.head.next
	var res []int
	for curr != nil {
		res = append(res, curr.val)
		curr = curr.next
	}
	return res
}
