func isValid(s string) bool {
	n := len(s)
	if n == 1 {
		return false
	}

	// dict
	dict := map[byte]byte{
		'[': ']',
		'(': ')',
		'{': '}',
	}

	// stack
	stack := make([]byte, 0)

	i := 0
	for i < n {
		// If current character is open bracket, push to stack
		if s[i] == '[' || s[i] == '(' || s[i] == '{' {
			stack = append(stack, s[i])
		}

		// If current character is close bracket, pop stack
		// check if this closed the latest open bracket
		if s[i] == ']' || s[i] == ')' || s[i] == '}' {
			// If stack empty, it fails
			if len(stack) == 0 {
				return false
			}

			// If stack is not empty, pop
			lastOpenBracket := stack[len(stack)-1]

			if dict[lastOpenBracket] != s[i] {
				return false
			}

			stack = stack[:len(stack)-1]
		}

		i++
	}

	if len(stack) > 0 {
		return false
	}

	return true
}