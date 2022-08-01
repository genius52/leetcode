package array

import (
	"sort"
)

func maximumGroups(grades []int) int {
	sort.Ints(grades)
	var l int = len(grades)
	var cur_len int = 1
	var groups int = 0
	var total int = 0
	for total < l{
		total += cur_len
		if total > l{
			break
		}
		groups++
		cur_len++
	}
	var sum int = 0
	for i := 0;i < l;i++{
		sum += grades[i]
	}
	for groups > 0{
		var pre_sum int = grades[0]
		var pre_len int = 1
		var cur_sum int = 0
		var cur_len int = 0
		var cur_groups int = 1
		for i := 1;i < l;i++{
			cur_sum += grades[i]
			cur_len++
			if cur_sum > pre_sum && cur_len > pre_len{
				pre_sum = cur_sum
				pre_len = cur_len
				cur_sum = 0
				cur_len = 0
				cur_groups++
			}
		}
		if cur_groups >= groups{
			return cur_groups
		}
		groups--
	}
	return groups
}