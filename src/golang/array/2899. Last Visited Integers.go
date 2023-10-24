package array

import "strconv"

func lastVisitedIntegers(words []string) []int {
	var l int = len(words)
	var res []int
	var nums []int
	var visit_idx int = 0
	for i := 0; i < l; i++ {
		if words[i] == "prev" {
			if visit_idx < 0 {
				res = append(res, -1)
			} else {
				res = append(res, nums[visit_idx])
				visit_idx--
			}
		} else {
			n, _ := strconv.Atoi(words[i])
			nums = append(nums, n)
			visit_idx = len(nums) - 1
		}
	}
	return res
}
