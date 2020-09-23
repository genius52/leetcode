package array

import "math"

//Input: prices = [1, 3, 2, 8, 4, 9], fee = 2
//Output: 8
//Explanation: The maximum profit can be achieved by:
//Buying at prices[0] = 1
//Selling at prices[3] = 8
//Buying at prices[4] = 4
//Selling at prices[5] = 9
//The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
func dp_maxProfit4(prices []int,pos int,fee int,has_stock bool,memo map[int]int)int{
	if pos >= len(prices){
		return 0
	}
	//if v,ok := memo[pos];ok{
	//	return v
	//}
	var profit int = 0
	if has_stock{
		sell := dp_maxProfit4(prices,pos + 1,fee,false,memo) + prices[pos]
		not_sell := dp_maxProfit4(prices,pos + 1,fee,true,memo)
		if sell > not_sell{
			profit = sell
		}else{
			profit = not_sell
		}
	}else{
		buy := dp_maxProfit4(prices,pos + 1,fee,true,memo) - prices[pos] - fee
		not_buy := dp_maxProfit4(prices,pos + 1,fee,false,memo)
		if buy > not_buy{
			profit = buy
		}else{
			profit = not_buy
		}
	}
	//memo[pos] = profit
	return profit
}

//func MaxProfit4(prices []int, fee int) int {
//	var memo map[int]int = make(map[int]int)
//	buy := dp_maxProfit4(prices,1,fee,true,memo) - prices[0] - fee
//	not_buy :=  dp_maxProfit4(prices,1,fee,false,memo)
//	if buy > not_buy{
//		return buy
//	}
//	return not_buy
//}

func MaxProfit4(prices []int, fee int) int {
	var l int = len(prices)
	var buy []int = make([]int,l)
	var sell []int = make([]int,l)
	var res int = 0
	buy[0] = -prices[0] - fee
	sell[0] = 0
	for i := 1;i < l;i++{
		buy[i] = int(math.Max(float64(sell[i - 1] - prices[i] - fee),float64(buy[i - 1])))
		sell[i] = int(math.Max(float64(buy[i - 1] + prices[i]),float64(sell[i - 1])))
		if buy[i] > res{
			res = buy[i]
		}
		if sell[i] > res{
			res = sell[i]
		}
	}
	return res
}