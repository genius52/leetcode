package array

import "sort"

func maxSum2815(nums []int) int {
	var l int = len(nums)
	var record [10][]int
	for i := 0; i < l; i++ {
		var max_num int = 0
		var n int = nums[i]
		for n > 0 {
			m := n % 10
			if m > max_num {
				max_num = m
				if max_num == 9 {
					break
				}
			}
			n /= 10
		}
		record[max_num] = append(record[max_num], nums[i])
	}
	var res int = -1
	for i := 1; i <= 9; i++ {
		var cur_len int = len(record[i])
		if cur_len <= 1 {
			continue
		}
		sort.Ints(record[i])
		sum := record[i][cur_len-1] + record[i][cur_len-2]
		if sum > res {
			res = sum
		}
	}
	return res
}
