package number

func MinOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	var l int = len(customers)
	var cur_profit int = 0
	var max_profit int = 0
	var max_profit_round = 0
	var people int = 0
	var idx int = 0
	var rounds int = 1
	for people > 0 || idx < l{
		if idx < l{
			people += customers[idx]
		}
		if people > 4{
			cur_profit += 4 * boardingCost - runningCost
			people -= 4
		}else{
			cur_profit += people * boardingCost - runningCost
			people = 0
		}
		if cur_profit > max_profit{
			max_profit = cur_profit
			max_profit_round = rounds
		}
		idx++
		rounds++
	}
	if max_profit <= 0{
		return -1
	}
	return max_profit_round
}