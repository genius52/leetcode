package array

import "math"

//Input: [1,3,5,4,7]
//Output: 2
//Explanation: The two longest increasing subsequence are [1, 3, 4, 7] and [1, 3, 5, 7].
//Input: [2,2,2,2,2]
//Output: 5
func FindNumberOfLIS(nums []int) int {
	var l int = len(nums)
	var lens []int = make([]int,l)//lens[i] store the length of element endup at nums[i]
	var counts []int = make([]int,l)//counts[i] store the count of subsequence end up at nums[i]
	var max_len int = 0
	for i := 0;i < l;i++{
		lens[i] = 1
		counts[i] = 1
		var cur_max_len int = 1
		for j := 0;j < i;j++{
			if nums[i] > nums[j]{
				cur_max_len = int(math.Max(float64(cur_max_len),float64(lens[j] + 1)))
				if lens[i] == lens[j] + 1{
					counts[i] += counts[j]
				}else if lens[i] < lens[j] + 1{
					lens[i] = lens[j] + 1
					counts[i] = counts[j]
				}
			}
		}
		if cur_max_len > max_len{
			max_len = cur_max_len
		}
	}
	var res int = 0
	for i := 0;i < l;i++{
		if lens[i] == max_len{
			res += counts[i]
		}
	}
	return res
}