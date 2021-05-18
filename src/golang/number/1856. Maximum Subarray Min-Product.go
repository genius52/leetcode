package number

import "container/list"

//Input: nums = [3,1,5,6,4,2]
//Output: 60
//Explanation: The maximum min-product is achieved with the subarray [5,6,4] (minimum value is 4).
//4 * (5+6+4) = 4 * 15 = 60.
//func cal_minproduct(nums []int,pos int)int{
//	var left int = pos - 1
//	var right int = pos + 1
//	var sum int64 = int64(nums[pos])
//	var l int = len(nums)
//	for left >= 0{
//		if nums[left] < nums[pos]{
//			break
//		}
//		sum += int64(nums[left])
//		sum = sum % 1000000007
//		left--
//	}
//	for right < l{
//		if nums[right] < nums[pos]{
//			break
//		}
//		sum += int64(nums[right])
//		sum = sum % 1000000007
//		right++
//	}
//	return int(int64(nums[pos]) * sum % 1000000007)
//}
//
//func maxSumMinProduct(nums []int) int {
//	var l int = len(nums)
//	var max_product int = 0
//	for i := 0;i < l;i++{
//		min_product := cal_minproduct(nums,i)
//		if min_product > max_product{
//			max_product = min_product
//		}
//	}
//	return max_product
//}

func MaxSumMinProduct(nums []int) int {
	var l int = len(nums)
	var pre_sum []int64 = make([]int64, l)
	pre_sum[0] = int64(nums[0])
	for i := 1; i < l; i++ {
		pre_sum[i] = pre_sum[i-1] + int64(nums[i])
	}
	var next_small [][2]int = make([][2]int,l)
	for i := 0;i < l;i++{
		next_small[i][0] = -1
		next_small[i][1] = -1
	}
	var q1 list.List
	for i := 0;i < l;i++{
		if q1.Len() == 0{
			q1.PushBack(i)
		}else{
			if nums[i] > nums[q1.Back().Value.(int)]{
				q1.PushBack(i)
			}else{
				for q1.Len() > 0 {
					var pre_index int = q1.Back().Value.(int)
					if nums[i] < nums[pre_index]{
						next_small[pre_index][1] = i
						q1.Remove(q1.Back())
					}else{
						break
					}
				}
				q1.PushBack(i)
			}
		}
	}
	var q2 list.List
	for i := l - 1;i >= 0;i--{
		if q2.Len() == 0{
			q2.PushBack(i)
		}else{
			if nums[i] > nums[q2.Back().Value.(int)]{
				q2.PushBack(i)
			}else{
				for q2.Len() > 0 {
					var pre_index int = q2.Back().Value.(int)
					if nums[i] < nums[pre_index]{
						next_small[pre_index][0] = i
						q2.Remove(q2.Back())
					}else{
						break
					}
				}
				q2.PushBack(i)
			}
		}
	}
	var max_product int64 = 0
	for i := 0;i < l;i++{
		var min_product int64 = 0
		if next_small[i][0] == -1 && next_small[i][1] == -1{//current nums is smallest
			min_product = int64(nums[i]) * pre_sum[l - 1]
		}else if next_small[i][0] == -1{//there is no smaller on left side
			min_product = int64(nums[i]) * pre_sum[next_small[i][1] - 1]
		}else if next_small[i][1] == -1{//there is no smaller on right side
			if i == 0{
				min_product = int64(nums[i]) * pre_sum[l - 1]
			}else{
				min_product = int64(nums[i]) * (pre_sum[l - 1] - pre_sum[next_small[i][0]])
			}
		}else{
			min_product = int64(nums[i]) * (pre_sum[next_small[i][1] - 1] - pre_sum[next_small[i][0]])
		}
		//min_product = min_product % 1000000007
		if max_product < min_product{
			max_product = min_product
		}
	}
	return int(max_product % 1000000007)
}