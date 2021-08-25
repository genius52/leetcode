package array

//Input: nums = [9,6,1,6,2]
//Output: 4
func MovesToMakeZigzag2(nums []int) int{
	var l int = len(nums)
	if l <= 1{
		return 0
	}
	var even_small int = 0
	var odd_small int = 0
	for i := 0;i < l;i++{
		if (i | 1) == i{//odd
			if i == l - 1{
				if nums[i] >= nums[i - 1]{
					odd_small += nums[i] - nums[i - 1] + 1
				}
			}else{
				if nums[i] >= nums[i - 1] || nums[i] >= nums[i + 1]{
					odd_small += nums[i] - min_int(nums[i - 1],nums[i + 1]) + 1
				}
			}
		}else{
			if i == 0{
				if nums[i] >= nums[i + 1]{
					even_small += nums[i] - nums[i + 1] + 1
				}
			}else if i == l - 1{
				if nums[i] >= nums[i - 1]{
					even_small += nums[i] - nums[i - 1] + 1
				}
			}else{
				if nums[i] >= nums[i - 1] || nums[i] >= nums[i + 1]{
					even_small += nums[i] - min_int(nums[i - 1],nums[i + 1]) + 1
				}
			}
		}
	}
	if even_small < odd_small{
		return even_small
	}else{
		return odd_small
	}
}

func MovesToMakeZigzag(nums []int) int {
	var l int = len(nums)
	if l <= 1{
		return 0
	}
	var change_even int = 0
	for i := 0;i < l;i += 2{ //even num should smaller
		if i == 0{
			if nums[i] >= nums[i + 1]{
				change_even = change_even + nums[i] - nums[i + 1] + 1
			}
		}else if i == l - 1{
			if nums[i] >= nums[i - 1]{
				change_even = change_even + nums[i] - nums[i - 1] + 1
			}
		} else {
			small_one := min_int(nums[i-1], nums[i+1])
			if nums[i] >= small_one {
				change_even = change_even + nums[i] - small_one + 1
			}
		}
	}
	var change_odd int = 0
	for i := 1;i < l;i += 2{ //even num should smaller
		if i == l - 1{
			if nums[i] >= nums[i - 1]{
				change_odd = change_odd + nums[i] - nums[i - 1] + 1
			}
		} else {
			small_one := min_int(nums[i-1], nums[i+1])
			if nums[i] >= small_one {
				change_odd = change_odd + nums[i] - small_one + 1
			}
		}
	}
	return min_int(change_even,change_odd)
}