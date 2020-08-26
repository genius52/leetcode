package number

import "sort"

func MaxCoins(piles []int) int {
	var l int = len(piles)
	var res int = 0
	sort.Ints(piles)
	begin := 0
	end := l - 2
	for begin <= end{
		res += piles[end]
		end -= 2
		begin++
	}
	return res
}