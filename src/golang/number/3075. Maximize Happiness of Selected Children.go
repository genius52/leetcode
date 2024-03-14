package number

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	var l int = len(happiness)
	var sum int64 = 0
	var rounds int = 0
	sort.Ints(happiness)
	for i := l - 1; i >= 0 && k > 0; i-- {
		if happiness[i] < rounds {
			break
		}
		sum += int64(happiness[i] - rounds)
		rounds++
		k--
	}
	return sum
}
