func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	// fmt.Println(s)
	re := regexp.MustCompile(`[^a-z0-9]+`)
	// fmt.Println(re)
	s = re.ReplaceAllString(s, "")
	// fmt.Println(s)

	// fmt.Println(s)
	n := len(s)

	left := 0
	right := n-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}