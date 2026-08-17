type Stack[T any] struct {
	items []T
}

// Push
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop
func (s *Stack[T]) Pop() {
	s.items = s.items[:s.Size()-1]
}

// Peek
func (s *Stack[T]) Peek() (T, bool) {
	size := s.Size()

	if size == 0 {
		var zero T
		return zero, false
	}

	top := s.items[size-1]

	return top, true
}

// Size
func (s *Stack[T]) Size() int {
	size := len(s.items)

	return size
}

var dict = map[byte]byte{
	'[': ']',
	'(': ')',
	'{': '}',
}

func isValid(s string) bool {
	// dict

	n := len(s)

	// Only 1 char. They don't have a pair
	// So it's invalid
	if n == 1 {
		return false
	}

	// Init stack
	stack := Stack[byte]{}

	i := 0
	for i < n {
		if _, exists := dict[s[i]]; exists {
			// If current character is open bracket, push to stack

			stack.Push(s[i])
		} else {
			// If current character is close bracket, peek and pop stack
			// Then check if this closed the latest open bracket

			// If stack empty, it fails
			if stack.Size() == 0 {
				return false
			}

			openBracket, _ := stack.Peek()
			stack.Pop()

			if dict[openBracket] != s[i] {
				return false
			}
		}

		i++
	}

	if stack.Size() > 0 {
		return false
	}

	return true
}
