package number

func MaxProfit1(prices []int) int {
	var l int = len(prices)
	var profit int = 0
	var low_price int = prices[0]
	for i := 1;i < l;i++{
		if prices[i] > low_price{
			if prices[i] - low_price > profit{
				profit = prices[i] - low_price
			}
		}else{
			low_price = prices[i]
		}
	}
	return profit
}

//func MaxProfit1(prices []int) int {
//	var l int = len(prices)
//	if l <= 1{
//		return 0
//	}
//	var res int = 0
//	var low_price int = prices[0]
//	for i := 1;i < l;i++{
//		if prices[i] > low_price{
//			profit := prices[i] - low_price
//			if profit > res{
//				res = profit
//			}
//		}else {
//			low_price = prices[i]
//		}
//	}
//	return res
//}