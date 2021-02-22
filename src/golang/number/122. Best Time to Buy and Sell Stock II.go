package number


//122
//func dp_profit(prices []int,start int,has_stock bool)int{
//	if start >= len(prices){
//		return 0
//	}
//	if has_stock{
//		return max_int(dp_profit(prices,start + 1,true),dp_profit(prices,start + 1,false) + prices[start])
//	}else{
//		return max_int(dp_profit(prices,start + 1,true) - prices[start],dp_profit(prices,start + 1,false))
//	}
//}
//
//func maxProfit(prices []int) int {
//	return dp_profit(prices,0,false)
//}
//
//func maxProfit2(prices []int)int{
//	var sum int = 0
//	for i := 1;i < len(prices);i++{
//		if prices[i] > prices[i-1]{
//			sum += prices[i] - prices[i-1]
//		}
//	}
//	return sum
//}

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