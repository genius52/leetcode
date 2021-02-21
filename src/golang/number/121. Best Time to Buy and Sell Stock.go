package number

//func dfs_maxProfit1(prices []int,index int,has_stock bool)int{
//	if index >= len(prices){
//		return 0
//	}
//	var profit int = 0
//	if has_stock{
//		not_sell := dfs_maxProfit1(prices,index + 1,has_stock)
//		sell := prices[index]
//		if not_sell > sell{
//			profit = not_sell
//		}else{
//			profit = sell
//		}
//	}else{
//		buy_now := dfs_maxProfit1(prices,index + 1,true) - prices[index]
//		not_buy := dfs_maxProfit1(prices,index + 1,false)
//		if buy_now > not_buy{
//			profit = buy_now
//		}else{
//			profit = not_buy
//		}
//	}
//	return profit
//}

func MaxProfit1(prices []int) int {
	var l int = len(prices)
	if l <= 1{
		return 0
	}
	var res int = 0
	var low_price int = prices[0]
	for i := 1;i < l;i++{
		if prices[i] > low_price{
			profit := prices[i] - low_price
			if profit > res{
				res = profit
			}
		}else {
			low_price = prices[i]
		}
	}
	return res
	//return dfs_maxProfit1(prices,0,false)
}