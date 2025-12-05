package leetcode

func MaxProfitTwo(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			dailyProfit := prices[i] - prices[i-1]

			maxProfit += dailyProfit
		}
	}

	return maxProfit
}