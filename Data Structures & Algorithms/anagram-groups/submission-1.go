func groupAnagrams(strs []string) [][]string {
    dict := make(map[string][]int)

    for idx, str := range strs {
        str = sortStr(str)

        dict[str] = append(dict[str], idx)
    }

    result := make([][]string, len(dict))
    resultIdx := 0
    for _, idxs := range dict {
        for _, idx := range idxs {
            result[resultIdx] = append(result[resultIdx], strs[idx])
        }

        resultIdx++
    }
    return result
}

// Bubble sort w/ flag
func sortStr(str string) string {
	chars := []byte(str)
	n := len(chars)

	for i := 0; i < n; i++ {
		swapped := false

		for j := 0; j < n-1-i; j++ {
			if chars[j] > chars[j+1] {
				chars[j], chars[j+1] = chars[j+1], chars[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return string(chars)
}
