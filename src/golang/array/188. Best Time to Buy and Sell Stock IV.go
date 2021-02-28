package array

import "strconv"

func dp_maxProfit188(k int,prices []int,l int,index int,has_stock bool,memo map[string]int)int{
	if index >= l{
		return 0
	}
	if k == 0 && !has_stock{
		return 0
	}
	var stock_tag int = 0
	if has_stock{
		stock_tag = 1
	}
	key := strconv.Itoa(index) + "," + strconv.Itoa(k) + "," + strconv.Itoa(stock_tag)
	if val,ok := memo[key];ok{
		return val
	}
	var max_profit int = 0
	if has_stock{
		//sell
		sellout := prices[index] + dp_maxProfit188(k,prices,l,index + 1,false,memo)
		//keep
		keep_stock := dp_maxProfit188(k,prices,l,index + 1,true,memo)
		max_profit = max_int(sellout,keep_stock)
	}else{
		//buy
		buy := -prices[index] + dp_maxProfit188(k - 1,prices,l,index + 1,true,memo)
		//not buy
		not_buy := dp_maxProfit188(k,prices,l,index + 1,false,memo)
		max_profit = max_int(buy,not_buy)
	}
	memo[key] = max_profit
	return max_profit
}

func MaxProfit(k int, prices []int) int {
	var l int = len(prices)
	var memo map[string]int = make(map[string]int) //index,k,has_stock
	return dp_maxProfit188(k,prices,l,0,false,memo)
}