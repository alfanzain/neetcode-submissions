func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	// fmt.Println(s)
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	// fmt.Println(re)
	s = re.ReplaceAllString(s, "")
	// fmt.Println(s)

	// fmt.Println(s)
	n := len(s)

	left := 0
	right := n-1
	for ; left < n/2 && right >= n/2; {
		sLeft := int(s[left])
		sRight := int(s[right])

		if sLeft < 97 {
			sLeft += 32
		}
		if sRight < 97 {
			sRight += 32
		}

		// fmt.Println("sLeft:", sLeft)
		// fmt.Println("sRight:", sRight)

		if sLeft != sRight {
			return false
		}
		left++
		right--
	}

	return true
}