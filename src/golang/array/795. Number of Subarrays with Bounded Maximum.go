package array

//Example:
//Input:
//nums = [2,1,4,3]
//left = 3
//right = 6

func NumSubarrayBoundedMax(nums []int, left int, right int) int{
	var l int = len(nums)
	var res int = 0
	var head int = -1//?
	var tail int = -1//?
	for i := 0;i < l;i++{
		if nums[i] > right{
			head = i
		}
		if nums[i] >= left{
			tail = i
		}
		res += tail - head
	}
	return res
}

//brute force
//func NumSubarrayBoundedMax(nums []int, left int, right int) int {
//	var l int = len(nums)
//	var res int = 0
//	for i := 0;i < l;i++{
//		if nums[i] > right{
//			continue
//		}
//		var cur_max int = nums[i]
//		for j := i;j < l;j++{
//			if nums[j] > right{
//				break
//			}else{
//				cur_max = max_int(cur_max,nums[j])
//				if cur_max >= left{
//					res++
//				}
//			}
//		}
//	}
//	return res
//}