package array

import "sort"

func longestSquareStreak(nums []int) int {
	var l int = len(nums)
	var record map[int]bool = make(map[int]bool)
	for _,n := range nums{
		record[n] = true
	}
	var visited map[int]bool = make(map[int]bool)
	sort.Ints(nums)
	var res int = -1
	for i := 0;i < l;i++{
		var n int = nums[i]
		if _,ok := visited[n];ok{
			continue
		}
		var cur_len int = 0
		for true{
			if _,ok := record[n * n];ok{
				cur_len++
				n = n * n
				visited[n] = true
			}else{
				break
			}
		}
		if cur_len > 0{
			if cur_len > res{
				res = cur_len + 1
			}
		}
	}
	return res
}