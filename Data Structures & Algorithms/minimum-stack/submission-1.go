type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() {
	s.items = s.items[:s.Size()-1]
}

func (s *Stack[T]) Peek() T {
	size := s.Size()

	var top T
	if size > 0 {
		top = s.items[size-1]
	}

	return top
}

func (s *Stack[T]) Size() int {
	size := len(s.items)
	return size

}

type MinStack struct {
	items Stack[int]
	minItems Stack[int]
}

func Constructor() MinStack {
	return MinStack{
		items: Stack[int]{},
		minItems: Stack[int]{},
	}
}

func (this *MinStack) Push(val int) {
	this.items.Push(val)

	// If there is no min, push to minItems
	if this.minItems.Size() == 0 {
		this.minItems.Push(val)
	} else {
		// Find the min
		min := this.minItems.Peek()
		if val < min {
			this.minItems.Push(val)
		} else {
			this.minItems.Push(min)
		}
	}
}

func (this *MinStack) Pop() {
	this.items.Pop()
	this.minItems.Pop()
}

func (this *MinStack) Top() int {
	return this.items.Peek()
}

func (this *MinStack) GetMin() int {
	return this.minItems.Peek()
}