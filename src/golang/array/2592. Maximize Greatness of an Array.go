package array

import "sort"

func MaximizeGreatness(nums []int) int {
	var l int = len(nums)
	//var nums2 []int = make([]int, l)
	//copy(nums2, nums)
	sort.Ints(nums)
	//var visited []bool = make([]bool, l)
	var res int = 0
	var left int = 0
	var right int = left + 1
	for left < l && right < l {
		for right < l && nums[left] >= nums[right] {
			right++
		}
		if right == l {
			break
		}
		//visited[right] = true
		right++
		res++
		left++
		//for left < l {
		//	left++
		//}
	}
	return res
}
