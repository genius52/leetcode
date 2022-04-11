package number

func ProductExceptSelf(nums []int) []int {
	var l int = len(nums)
	var left []int = make([]int,l)
	var right []int = make([]int,l)
	left[0] = nums[0]
	right[l - 1] = nums[l - 1]
	for i := 1;i < l;i++{
		left[i] = left[i - 1] * nums[i]
		right[l - 1 - i] = right[l - i] * nums[l - 1 - i]
	}
	var res []int = make([]int,l)
	for i := 0;i < l;i++{
		if i == 0{
			res[i] = right[i + 1]
		}else if i == l - 1{
			res[i] = left[i - 1]
		}else{
			res[i] = left[i - 1] * right[i + 1]
		}
	}
	return res
}

func productExceptSelf(nums []int) []int{
	var l int = len(nums)
	var res []int = make([]int,l)
	var left_sum int = 1
	var left_zero_idx int = -1
	for i := 0;i < l;i++{
		if nums[i] == 0{
			left_zero_idx = i
			break
		}
		left_sum *= nums[i]
	}
	var right_sum int = 1
	var right_zero_idx int = -1
	for i := l - 1;i >= 0;i--{
		if nums[i] == 0{
			right_zero_idx = i
			break
		}
		right_sum *= nums[i]
	}
	if left_zero_idx == -1{
		for i := 0;i < l;i++{
			res[i] = left_sum / nums[i]
		}
	}else{
		if left_zero_idx == right_zero_idx{
			res[left_zero_idx] = left_sum * right_sum
		}
	}
	return res
}