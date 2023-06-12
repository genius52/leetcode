package number

import "sort"

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	rest := money - prices[0] - prices[1]
	if rest < 0 {
		return money
	}
	return rest
}
