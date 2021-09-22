package array

func maxValueAfterReverse(nums []int) int {
	var l int = len(nums)
	var sum int = 0
	for i := 1;i < l;i++{
		sum += abs_int(nums[i] - nums[i - 1])
	}
	var prefix_max []int = make([]int,l)
	prefix_max[0] = 0
	var postfix_max []int = make([]int,l)
	for i := 1;i < l;i++{
		prefix_max[i] = max_int(prefix_max[i - 1],abs_int(nums[i] - nums[i - 1]))
	}
	for i := l - 2;i >= 0;i--{
		postfix_max[i] = max_int(postfix_max[i + 1],abs_int(nums[i] - nums[i + 1]))
	}
	var max_profit int = 0
	for i := 0;i < l;i++{

	}
	for i := 1;i < l - 1;i++{

	}
	return sum + max_profit
}

//TLE
func MaxValueAfterReverse(nums []int) int {
	var l int = len(nums)
	var sum int = 0
	for i := 1;i < l;i++{
		sum += abs_int(nums[i] - nums[i - 1])
	}
	var max_profit int = 0
	for i := 0;i < l;i++{
		for j := i + 1;j < l;j++{
			if i == 0 && j == l - 1{
				continue
			}
			//swap nums[i] ... nums[j]
			if i == 0{
				cur := abs_int(nums[j + 1] - nums[i]) - abs_int(nums[j + 1] - nums[j])
				if cur > max_profit{
					max_profit = cur
				}
			}else if j == (l - 1){
				cur := abs_int(nums[i - 1] - nums[j]) - abs_int(nums[i - 1] - nums[i])
				if cur > max_profit{
					max_profit = cur
				}
			}else{
				cur := abs_int(nums[i - 1] - nums[j]) + abs_int(nums[j + 1] - nums[i]) -
					(abs_int(nums[i - 1] - nums[i]) + abs_int(nums[j + 1] - nums[j]))
				if cur > max_profit{
					max_profit = cur
				}
			}
		}
	}
	return sum + max_profit
}