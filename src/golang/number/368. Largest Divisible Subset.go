package number

import "sort"

//Input: nums = [1,2,3]
//Output: [1,2]
func LargestDivisibleSubset(nums []int) []int {
	var l int = len(nums)
	if l == 1{
		return nums
	}
	sort.Ints(nums)
	var dp [][]int = make([][]int,l)
	dp[0] = []int{nums[0]}
	var res []int
	var max_len int = 0
	for i := 1;i < l;i++{
		var cur_data []int
		for j := i - 1;j >= 0;j--{
			if nums[i] % nums[j] != 0{
				continue
			}
			if len(dp[j]) >= len(cur_data){
				cur_data = dp[j]
			}
		}
		var add []int = make([]int,len(cur_data))
		copy(add,cur_data)
		add = append(add,nums[i])
		var cur_len int = len(add)
		if cur_len > max_len{
			max_len = cur_len
			res = add
		}
		dp[i] = add
	}
	return res
}