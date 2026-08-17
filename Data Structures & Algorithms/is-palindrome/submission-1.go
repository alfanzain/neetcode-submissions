func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	
	s = strings.ReplaceAll(s, " ", "")
	n := len(s)

	left := 0
	right := n-1
	for ; left < n/2 && right >= n/2; {
		// Skip non alphanumeric
		if !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
			continue
		}
		if !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
			continue
		}

		sLeft := int(s[left])
		sRight := int(s[right])

		if sLeft < 97 {
			sLeft += 32
		}
		if sRight < 97 {
			sRight += 32
		}

		fmt.Println("sLeft:", sLeft)
		fmt.Println("sRight:", sRight)

		if sLeft != sRight {
			return false
		}
		left++
		right--
	}
	
	return true
}
