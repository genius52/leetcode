package array

import "sort"

func MaxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	var l int = len(nums)
	var res int = 0
	var left int = (l - 1) / 2
	var right int = l - 1
	var visited []bool = make([]bool, l)
	for left >= 0 && right >= 0 {
		for left >= 0 && nums[left]*2 > nums[right] && !visited[left] {
			left--
		}
		if left < 0 {
			break
		}
		visited[left] = true
		visited[right] = true
		res += 2
		for right >= 0 && visited[right] {
			right--
		}
		left--
	}
	return res
}
