package array

func sumOfBeauties(nums []int) int {
	var l int = len(nums)
	if l <= 2{
		return 0
	}
	var prefix_max []int = make([]int,l)
	prefix_max[0] = nums[0]
	var postfix_min []int = make([]int,l)
	postfix_min[l - 1] = nums[l - 1]
	for i := 1;i < l;i++{
		prefix_max[i] = max_int(prefix_max[i - 1],nums[i])
	}
	for i := l - 2;i >= 0;i--{
		postfix_min[i] = min_int(postfix_min[i + 1],nums[i])
	}
	var res int = 0
	for i := 1;i <= l - 2;i++{
		if prefix_max[i - 1] < nums[i] && nums[i] < postfix_min[i + 1]{
			res += 2
		}else if nums[i - 1] < nums[i] && nums[i] < nums[i + 1]{
			res++
		}
	}
	return res
}