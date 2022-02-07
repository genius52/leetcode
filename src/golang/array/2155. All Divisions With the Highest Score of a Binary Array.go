package array

func maxScoreIndices(nums []int) []int {
	var l int = len(nums)
	var prefix_zero []int = make([]int,l)
	var suffix_one []int = make([]int,l)
	if nums[0] == 0{
		prefix_zero[0] = 1
	}
	for i := 1;i < l;i++{
		prefix_zero[i] = prefix_zero[i - 1]
		if nums[i] == 0{
			prefix_zero[i]++
		}
	}
	if nums[l - 1] == 1{
		suffix_one[l - 1] = 1
	}
	for i := l - 2;i >= 0;i--{
		suffix_one[i] = suffix_one[i + 1]
		if nums[i] == 1{
			suffix_one[i]++
		}
	}
	var max_val int = prefix_zero[l - 1]
	var res []int = []int{l}
	if suffix_one[0] > max_val{
		res = []int{0}
		max_val = suffix_one[0]
	}else if suffix_one[0] == max_val{
		res = append(res,0)
	}
	for i := 0;i < l - 1;i++{
		cur := prefix_zero[i] + suffix_one[i + 1]
		if cur > max_val{
			max_val = cur
			res = []int{i + 1}
		}else if cur == max_val{
			res = append(res,i + 1)
		}
	}
	return res
}