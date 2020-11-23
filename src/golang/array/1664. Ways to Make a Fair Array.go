package array

//Input: nums = [2,1,6,4]
//Output: 1
//Explanation:
//Remove index 0: [1,6,4] -> Even sum: 1 + 4 = 5. Odd sum: 6. Not fair.
//Remove index 1: [2,6,4] -> Even sum: 2 + 4 = 6. Odd sum: 6. Fair.
//Remove index 2: [2,1,4] -> Even sum: 2 + 4 = 6. Odd sum: 1. Not fair.
//Remove index 3: [2,1,6] -> Even sum: 2 + 6 = 8. Odd sum: 1. Not fair.
//There is 1 index that you can remove to make nums fair.
func WaysToMakeFair(nums []int) int {
	var l int = len(nums)
	if l == 1{
		return 1
	}
	var left_even []int = make([]int,l)//i左边的偶数和
	var left_odd []int = make([]int,l)//i左边的奇数和
	left_even[0] = nums[0]
	left_odd[1] = nums[1]
	var right_even []int = make([]int,l)//i右边的偶数和
	var right_odd []int = make([]int,l)//i右边的奇数和
	if l | 1 == l{
		right_even[l - 1] = nums[l - 1]
		right_odd[l - 2] = nums[l - 2]
	}else{
		right_odd[l - 1] = nums[l - 1]
		right_even[l - 2] = nums[l - 2]
	}
	var total int = nums[0] + nums[1]
	for i := 2;i < l;i++{
		if i | 1 == i{
			left_odd[i] = left_odd[i - 2] + nums[i]
		}else{
			left_even[i] = left_even[i - 2] + nums[i]
		}
		total += nums[i]
	}
	for i := l - 3;i >= 0;i--{
		if i | 1 == i{
			right_odd[i] = right_odd[i + 2] + nums[i]
		}else{
			right_even[i] = right_even[i + 2] + nums[i]
		}
	}
	var res int = 0
	for i := 0;i < l;i++{
		var cur_leftodd int = 0
		var cur_lefteven int = 0
		var cur_rightodd int = 0
		var cur_righteven int = 0
		var all_odd int = 0
		var all_even int = 0
		if i == 0{
			cur_lefteven = 0
			cur_leftodd = 0
		}else if i == 1{
			cur_leftodd = 0
			cur_lefteven = left_even[0]
		}else{
			if i | 1 == i{
				cur_lefteven = left_even[i - 1]
				cur_leftodd = left_odd[i - 2]
			}else{
				cur_lefteven = left_even[i - 2]
				cur_leftodd = left_odd[i - 1]
			}
		}

		if i == l - 1{
			cur_rightodd = 0
			cur_righteven = 0
		}else if i == l - 2{
			if i | 1 == i{
				cur_rightodd = right_even[i + 1]
				cur_righteven = 0
			}else{
				cur_righteven = right_odd[i + 1]
				cur_rightodd = 0
			}
		}else{
			if i | 1 == i{
				cur_rightodd = right_even[i + 1]
				cur_righteven = right_odd[i + 2]
			}else{
				cur_rightodd = right_even[i + 2]
				cur_righteven = right_odd[i + 1]
			}
		}
		all_odd = cur_leftodd + cur_rightodd
		all_even = cur_lefteven + cur_righteven
		if all_odd == all_even{
			res++
		}
	}
	return res
}