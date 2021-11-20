package array

import "sort"

func MaxProfit1648(inventory []int, orders int) int {
	sort.Ints(inventory)
	var l int = len(inventory)
	var profit int = 0
	var dup_cnt int = 1
	for i := l - 1;i > 0;i--{
		cur := inventory[i]
		next := inventory[i - 1]
		diff := (cur - next) * dup_cnt
		if orders > diff {
			profit += (next + 1 + cur) * diff/2
			profit %= (1e9 + 7)
			orders -= diff
			dup_cnt++
		}else{
			for orders >= dup_cnt{
				profit += cur * dup_cnt
				cur--
				orders -= dup_cnt
			}
			profit %= 1e9 + 7
			if orders > 0{
				profit += cur * orders
				profit %= 1e9 + 7
				return profit
			}
		}
	}
	cur := inventory[0]
	for orders >= dup_cnt{
		profit += cur * dup_cnt
		cur--
		orders -= dup_cnt
	}
	profit %= 1e9 + 7
	if orders > 0{
		profit += cur * orders
		profit %= 1e9 + 7
	}
	return profit
}