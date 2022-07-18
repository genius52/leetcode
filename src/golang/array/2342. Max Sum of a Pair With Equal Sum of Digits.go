package array

import "sort"

func maximumSum2(nums []int) int {
	var l int = len(nums)
	var record map[int][]int = make(map[int][]int)
	for i := 0;i < l;i++{
		n := nums[i]
		var sum int = 0
		for n > 0{
			sum += n % 10
			n /= 10
		}
		record[sum] = append(record[sum],nums[i])
	}
	var res int = -1
	for _,idx := range record{
		var cur_len int = len(idx)
		if cur_len <= 1{
			continue
		}
		sort.Ints(idx)
		cur := idx[cur_len - 1] + idx[cur_len - 2]
		if cur > res{
			res = cur
		}
	}
	return res
}