package array

import "sort"

func FindScore(nums []int) int64 {
	var l int = len(nums)
	var number_oldidx [][2]int = make([][2]int, l)
	for i := 0; i < l; i++ {
		number_oldidx[i] = [2]int{nums[i], i}
	}
	sort.Slice(number_oldidx, func(i, j int) bool {
		if number_oldidx[i][0] != number_oldidx[j][0] {
			return number_oldidx[i][0] < number_oldidx[j][0]
		} else {
			return number_oldidx[i][1] < number_oldidx[j][1]
		}
	})
	var oldidx_newidx []int = make([]int, l)
	for i := 0; i < l; i++ {
		old_idx := number_oldidx[i][1]
		oldidx_newidx[old_idx] = i
	}
	var res int64 = 0
	var visited []bool = make([]bool, l)
	for i := 0; i < l; i++ {
		if visited[number_oldidx[i][1]] {
			continue
		}
		res += int64(number_oldidx[i][0])
		old_idx := number_oldidx[i][1]
		var left_idx int = old_idx - 1
		var right_idx int = old_idx + 1
		for left_idx >= 0 && visited[left_idx] {
			left_idx--
		}
		if left_idx >= 0 {
			visited[left_idx] = true
		}
		for right_idx < l && visited[right_idx] {
			right_idx++
		}
		if right_idx < l {
			visited[right_idx] = true
		}
	}
	return res
}
