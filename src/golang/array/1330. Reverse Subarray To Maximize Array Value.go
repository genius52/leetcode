package array

func MaxValueAfterReverse(nums []int) int {
	var l int = len(nums)
	var sum int = 0
	for i := 1;i < l;i++{
		sum += abs_int(nums[i] - nums[i - 1])
	}
	var prefix_max_diff []int = make([]int,l)
	prefix_max_diff[1] = abs_int(nums[1] - nums[0])
	var prefix_min_diff []int = make([]int,l)
	prefix_min_diff[1] = abs_int(nums[1] - nums[0])
	var suffix_max_diff []int = make([]int,l)
	suffix_max_diff[l - 2] = abs_int(nums[l - 1] - nums[l - 2])
	var suffix_min_diff []int = make([]int,l)
	suffix_min_diff[l - 2] = abs_int(nums[l - 1] - nums[l - 2])

	for i := 2;i < l;i++{
		prefix_max_diff[i] = max_int(prefix_max_diff[i - 1],abs_int(nums[i] - nums[i - 1]))
		prefix_min_diff[i] = min_int(prefix_min_diff[i - 1],abs_int(nums[i] - nums[i - 1]))
	}
	for i := l - 3;i >= 0;i--{
		suffix_max_diff[i] = max_int(suffix_max_diff[i + 1],abs_int(nums[i] - nums[i + 1]))
		suffix_min_diff[i] = min_int(suffix_min_diff[i + 1],abs_int(nums[i] - nums[i + 1]))
	}
	var max_profit int = 0
	for i := 1;i < l - 1;i++{
		max_profit = max_int(max_profit,prefix_max_diff[i] + suffix_max_diff[i + 1])
	}
	return sum + max_profit
}

//TLE
//func MaxValueAfterReverse(nums []int) int {
//	var l int = len(nums)
//	var sum int = 0
//	for i := 1;i < l;i++{
//		sum += abs_int(nums[i] - nums[i - 1])
//	}
//	var max_profit int = 0
//	for i := 0;i < l;i++{
//		for j := i + 1;j < l;j++{
//			if i == 0 && j == l - 1{
//				continue
//			}
//			//swap nums[i] ... nums[j]
//			if i == 0{
//				cur := abs_int(nums[j + 1] - nums[i]) - abs_int(nums[j + 1] - nums[j])
//				if cur > max_profit{
//					max_profit = cur
//				}
//			}else if j == (l - 1){
//				cur := abs_int(nums[i - 1] - nums[j]) - abs_int(nums[i - 1] - nums[i])
//				if cur > max_profit{
//					max_profit = cur
//				}
//			}else{
//				cur := abs_int(nums[i - 1] - nums[j]) + abs_int(nums[j + 1] - nums[i]) -
//					(abs_int(nums[i - 1] - nums[i]) + abs_int(nums[j + 1] - nums[j]))
//				if cur > max_profit{
//					max_profit = cur
//				}
//			}
//		}
//	}
//	return sum + max_profit
//}