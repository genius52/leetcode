package number

import "sort"

func MaxRunTime(n int, batteries []int) int64 {
	var l int = len(batteries)
	if l < n{
		return 0
	}
	sort.Ints(batteries)
	var sum int64 = 0
	for i := 0;i < l;i++{
		sum += int64(batteries[i])
	}
	var left int64 = 0
	var right int64 = sum / int64(n)
	var res int64 = left
	for left < right{
		var target_hours int64 = (left + right + 1)/2
		var cur_sum int64 = 0
		for i := l - 1;i >= 0;i--{
			if int64(batteries[i]) > target_hours{
				cur_sum += target_hours
			}else{
				cur_sum += int64(batteries[i])
			}
		}
		if int64(n) * target_hours <= cur_sum{
			left = target_hours
			res = left
		}else{
			right = target_hours - 1
			res = right
		}
	}
	return res
}