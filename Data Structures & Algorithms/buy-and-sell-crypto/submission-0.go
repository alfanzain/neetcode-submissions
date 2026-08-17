func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	n := len(prices)
	left, right := n-2, n-1
	maxProfit := 0
	for right > 0 {
		// right price - left price = profit
		profit := prices[right] - prices[left]
		// fmt.Println("profit:", profit)

		if maxProfit < profit {
			maxProfit = profit
		}
		
		if left == 0 {
			right--
			left = right - 1
		} else {
			left--
		}
	}

	return maxProfit
}