package array

func MinimumMountainRemovals(nums []int) int {
	var l int = len(nums)
	var increase []int = make([]int,l)//increase[i]: the longest length from 0 to i by increase order
	var decrease []int = make([]int,l)//decrease[i]: the longest length from l - i to i by decrease order
	for i := 0;i < l;i++{
		increase[i] = 1
		decrease[i] = 1
	}
	for i := 0;i < l;i++{
		for j := i - 1;j >= 0;j--{
			if nums[i] > nums[j]{
				increase[i] = max_int(increase[i],1 + increase[j])
			}
		}
	}
	for i := l - 1;i >= 0;i--{
		for j := i + 1;j < l;j++{
			if nums[i] > nums[j]{
				decrease[i] = max_int(decrease[i],1 + decrease[j])
			}
		}
	}
	var res int = 0
	for i := 1;i < l - 1;i++{
		if increase[i] > 1 && decrease[i] > 1{
			cur_len := increase[i] + decrease[i] - 1
			if cur_len > res{
				res = cur_len
			}
		}
	}
	return l - res
}