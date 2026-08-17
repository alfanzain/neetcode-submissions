func maxProfit(prices []int) int {
	n := len(prices)
	if n == 1 {
		return 0
	}

	buy := 0
	sell := 1
	maxProfit := 0

	for sell < n {
		fmt.Println("buy:", buy, "sell:", sell)
		profit := prices[sell] - prices[buy]
		if profit > maxProfit {
			maxProfit = profit
		}
		if prices[sell] < prices[buy] {
			buy = sell
		}
		sell++
	}

	return maxProfit
}