package number

//Input: prices = [3,3,5,0,0,3,1,4]
//Output: 6
//Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
//Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
func dp_maxProfit(prices []int,l int,index int,has_stock bool,allowed int,memo [][3][2]int)int{
	if index >= l{
		return 0
	}
	if allowed == 0 && !has_stock{
		return 0
	}
	var status int = 0
	if has_stock{
		status = 1
	}
	if memo[index][allowed][status] != -1{
		return memo[index][allowed][status]
	}
	var max_profit int = 0
	if has_stock{
		sell_now := prices[index] + dp_maxProfit(prices,l,index + 1,false,allowed,memo)
		not_sell := dp_maxProfit(prices,l,index + 1,true,allowed,memo)
		if sell_now > not_sell{
			max_profit = sell_now
		}else{
			max_profit = not_sell
		}
	}else{
		var buy_now int = 0
		if allowed > 0{
			buy_now = -prices[index] + dp_maxProfit(prices,l,index + 1,true,allowed - 1,memo)
		}
		not_buy := dp_maxProfit(prices,l,index + 1,false,allowed,memo)
		if buy_now > not_buy{
			max_profit = buy_now
		}else{
			max_profit = not_buy
		}
	}
	memo[index][allowed][status] = max_profit
	return max_profit
}

func MaxProfit(prices []int) int {
	var l int = len(prices)
	var memo [][3][2]int = make([][3][2]int,l)
	for i := 0;i < l;i++{
		memo[i][0][0] = -1
		memo[i][0][1] = -1
		memo[i][1][0] = -1
		memo[i][1][1] = -1
		memo[i][2][0] = -1
		memo[i][2][1] = -1
	}
	return dp_maxProfit(prices,l,0,false,2,memo)
}