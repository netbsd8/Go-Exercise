package main

func maxProfit2(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}

	ans := 0
	curMin := prices[0]
	for i:=0; i<length; i++ {
		if prices[i] < curMin {
			curMin = prices[i]
		} else {
			ans = ans + (prices[i] - curMin)
			curMin = prices[i]
		}
	}

	return ans
}

func maxProfit22(prices []int) int {
	length := len(prices)
	ans := 0
	for i:=0; i<length-1; i++ {
		if prices[i] < prices[i+1] {
			ans = ans + (prices[i+1] - prices[i])
		}
	}
	return ans
}