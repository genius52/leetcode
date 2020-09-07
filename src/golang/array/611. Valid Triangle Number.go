package array

import "sort"

//Input: [2,2,3,4]
//Output: 3
//Explanation:
//Valid combinations are:
//2,3,4 (using the first 2)
//2,3,4 (using the second 2)
//2,2,3
func TriangleNumber(nums []int) int {
	var l int = len(nums)
	sort.Ints(nums)
	var res int = 0
	for big := l - 1;big >= 2;big--{
		small := 0
		mid := big - 1
		for small < mid{
			if nums[small] + nums[mid] > nums[big]{
				res += mid - small
				mid--
			}else{
				small++
			}
		}
	}
	return res
}