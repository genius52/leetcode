package array

import "sort"

func check_valid(nums []int)(bool){
	var l int = len(nums)
	var new_nums []int = make([]int,l)
	copy(new_nums,nums)
	sort.Ints(new_nums)
	for i := 1;i < l - 1;i++{
		if new_nums[i] - new_nums[i - 1] != new_nums[i + 1] - new_nums[i]{
			return false
		}
	}
	return true
}

func CheckArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var steps int = len(l)
	var res []bool = make([]bool,steps)
	for i := 0;i < steps;i++{
		begin := l[i]
		end := r[i]
		res[i] = check_valid(nums[begin:end + 1])
	}
	return res
}