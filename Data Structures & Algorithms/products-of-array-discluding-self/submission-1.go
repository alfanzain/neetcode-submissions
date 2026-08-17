func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	zeroCount := 0
	zeroLastIndex := 0
	total := 1

	for i, num := range nums {
		if num != 0 {
			total *= num
		} else {
			zeroCount++

			if zeroCount > 1 {
				return result
			}

			zeroLastIndex = i
		}
	}

	if zeroCount == 1 {
		result[zeroLastIndex] = total
		return result
	}

	for i, num := range nums {
		result[i] = total / num
	}

	return result
}