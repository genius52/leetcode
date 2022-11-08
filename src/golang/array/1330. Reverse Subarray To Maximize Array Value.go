package array

func MaxValueAfterReverse(nums []int) int {
	var l int = len(nums)
	var sum int = 0
	for i := 1;i < l;i++{
		sum += abs_int(nums[i] - nums[i - 1])
	}
	var max_profit int = 0
	//减少的值 = |nums[i] - nums[i + 1]| + |nums[j] - nums[j - 1]|
	//增加的值 = |nums[i] - nums[j - 1]| + |nums[j] - nums[i + 1]|
	//变化的值 = |nums[i] - nums[j - 1]| + |nums[j] - nums[i + 1]| - ( |nums[i] - nums[i + 1]| + |nums[j] - nums[j - 1]| )
	for i := 1;i < l - 1;i++{

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