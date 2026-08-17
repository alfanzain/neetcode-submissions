func productExceptSelf(nums []int) []int {
	n := len(nums)
	zeroCount := 0
	total := 1
	for _, num := range nums {
		if num != 0 {
			total *= num
		} else {
			zeroCount++
		}
	}

	if zeroCount > 1 {
		return make([]int, n)
	}

	result := make([]int, n)
	for i, num := range nums {
		if num != 0 && zeroCount == 0 {
			result[i] = total / num
		}
		if num != 0 && zeroCount > 0 {
			result[i] = 0
		}
		if num == 0 {
			result[i] = total	
		}
	}

	return result
}
