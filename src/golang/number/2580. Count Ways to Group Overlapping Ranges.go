package number

import (
	"sort"
)

func countWays(ranges [][]int) int {
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] == ranges[j][0] {
			return ranges[i][1] < ranges[j][1]
		}
		return ranges[i][0] < ranges[j][0]
	})
	var l int = len(ranges)
	var pre_end int = ranges[0][1]
	var groups int = 1
	for i := 1; i < l; i++ {
		if ranges[i][0] > pre_end {
			pre_end = ranges[i][1]
			groups++
		} else {
			pre_end = max_int(pre_end, ranges[i][1])
		}
	}
	var res int = 1
	var MOD int = 1e9 + 7
	for groups > 0 {
		res *= 2
		res %= MOD
		groups--
	}
	return res
}
