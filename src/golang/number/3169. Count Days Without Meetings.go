package number

import "sort"

func CountDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][0] != meetings[j][0] {
			return meetings[i][0] < meetings[j][0]
		} else {
			return meetings[i][1] > meetings[j][1]
		}
	})
	var res int = 0
	var pre_end int = 0
	for i := 0; i < len(meetings); i++ {
		if meetings[i][0] > pre_end {
			res += meetings[i][0] - pre_end - 1
			pre_end = meetings[i][1]
		} else {
			pre_end = max_int(pre_end, meetings[i][1])
		}
	}
	res += days - pre_end
	return res
}
