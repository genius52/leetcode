package number

import "math"

//Input: [1,2,3,0,2]
//Output: 3
//Explanation: transactions = [buy, sell, cooldown, buy, sell]
func dp_maxProfit3(prices []int,cur_pos int,has_stock bool,is_cooldown bool,nostock_memo map[int]int,hasstock_memo map[int]int)int{
	if cur_pos >= len(prices){
		return 0
	}

	var max int = math.MinInt32
	if is_cooldown{
		max = dp_maxProfit3(prices,cur_pos + 1,has_stock,false,nostock_memo,hasstock_memo)
	}else{
		if has_stock{
			if _,ok := hasstock_memo[cur_pos];ok{
				return hasstock_memo[cur_pos]
			}
			max = max_int(prices[cur_pos] + dp_maxProfit3(prices,cur_pos + 1,false,true,nostock_memo,hasstock_memo),
				dp_maxProfit3(prices,cur_pos + 1,true,false,nostock_memo,hasstock_memo))
			hasstock_memo[cur_pos] = max
		}else{
			if _,ok := nostock_memo[cur_pos];ok{
				return nostock_memo[cur_pos]
			}
			max = max_int(dp_maxProfit3(prices,cur_pos + 1,true,false,nostock_memo,hasstock_memo) - prices[cur_pos],
				dp_maxProfit3(prices,cur_pos + 1,false,false,nostock_memo,hasstock_memo))
			nostock_memo[cur_pos] = max
		}
	}
	return max
}

func maxProfit3(prices []int) int {
	var nostock_record map[int]int = make(map[int]int)
	var hasstock_record map[int]int = make(map[int]int)
	return dp_maxProfit3(prices,0,false,false,nostock_record,hasstock_record)
}