package array

import "sort"

func matrixSum(nums [][]int) int {
	var rows int = len(nums)
	var columns int = len(nums[0])
	for i := 0; i < rows; i++ {
		sort.Ints(nums[i])
	}
	var res int = 0
	for j := 0; j < columns; j++ {
		var max_score int = -1
		for i := 0; i < rows; i++ {
			if nums[i][j] > max_score {
				max_score = nums[i][j]
			}
		}
		res += max_score
	}
	return res
}
