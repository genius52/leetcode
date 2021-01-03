package array

//Input: nums = [1,2,3,4,4]
//Output: 12
func MaxSumDivThree(nums []int) int {
	var l int = len(nums)
	var mod_zero int = 0
	var mod_one int = 0
	var mod_two int = 0
	m := nums[0] % 3
	if m == 0{
		mod_zero = nums[0]
	}else if m == 1{
		mod_one = nums[0]
	}else if m == 2{
		mod_two = nums[0]
	}
	for i := 1;i < l;i++{
		mod := nums[i] % 3
		var new_zero,new_one,new_two int = 0,0,0
		if mod == 0{
			mod_zero += nums[i]
			if mod_one > 0{
				mod_one += nums[i]
			}
			if mod_two > 0{
				mod_two += nums[i]
			}
		}else if mod == 1{
			if mod_two > 0{
				new_zero = max_int(mod_zero,mod_two + nums[i])
			}
			if mod_zero > 0{
				new_one = max_int(mod_one,mod_zero + nums[i])
			}else{
				new_one = nums[i]
			}
			if mod_one > 0 {
				new_two = max_int(mod_two, mod_one+nums[i])
			}
		}else if mod == 2{
			if mod_one > 0{
				new_zero = max_int(mod_zero,mod_one + nums[i])
			}
			if mod_two > 0{
				new_one = max_int(mod_one,mod_two + nums[i])
			}
			if mod_zero > 0{
				new_two = max_int(mod_two,mod_zero + nums[i])
			}else{
				new_two = nums[i]
			}
		}
		mod_zero = max_int(mod_zero,new_zero)
		mod_one = max_int(mod_one,new_one)
		mod_two = max_int(mod_two,new_two)
	}
	return mod_zero
}