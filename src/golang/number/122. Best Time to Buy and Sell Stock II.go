package number

func MaxProfit2(prices []int) int{
	var l int = len(prices)
	var res int = 0
	for i := 1;i < l;i++{
		if prices[i] > prices[i - 1]{
			res += prices[i] - prices[i-1]
		}
	}
	return res
}