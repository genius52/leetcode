package array

//Input: nums = [4,2,4,5,6]
//Output: 17

func MaximumUniqueSubarray(nums []int) int{
	var l int = len(nums)
	var left int = 0
	var right int = 0
	var record map[int]bool = make(map[int]bool)
	var cur_sum int = 0
	var res int = 0
	for left < l{
		for right < l{
			if _,ok := record[nums[right]];ok{
				break
			}
			cur_sum += nums[right]
			record[nums[right]] = true
			right++
			res = max_int(res,cur_sum)
		}
		cur_sum -= nums[left]
		delete(record,nums[left])
		left++
	}
	return res
}

//Explanation: The optimal subarray here is [2,4,5,6].
//func MaximumUniqueSubarray(nums []int) int{
//	var l int = len(nums)
//	var res int = 0
//	var record [10001]int
//	for i := 0;i <= 10000;i++{
//		record[i] = -1
//	}
//	var prefix []int = make([]int,l + 1)
//	for i := 0;i < l;i++{
//		prefix[i + 1] = nums[i] + prefix[i]
//	}
//	var left int = 0
//	var right int = 0
//	for left < l{
//		for right < l && record[nums[right]] < left{
//			record[nums[right]] = right
//			res = max_int(res,prefix[right + 1] - prefix[left])
//			right++
//		}
//		if right >= l{
//			break
//		}
//		left = record[nums[right]] + 1
//		record[nums[right]] = right
//		right++
//	}
//	return res
//}

//func MaximumUniqueSubarray(nums []int) int {
//	var l int = len(nums)
//	var begin int = 0
//	var end int = 0
//	var record [10001]int
//	var dup_map map[int]bool = make(map[int]bool)
//	var cur_sum int = 0
//	var max_sum int = 0
//	for end < l{
//		cur_sum += nums[end]
//		record[nums[end]]++
//		if record[nums[end]] > 1{
//			dup_map[nums[end]] = true
//		}
//		if len(dup_map) == 0{
//			if cur_sum > max_sum{
//				max_sum = cur_sum
//			}
//		}else{
//			for len(dup_map) > 0 && begin < end{
//				cur_sum -= nums[begin]
//				record[nums[begin]]--
//				if record[nums[begin]] == 1{
//					delete(dup_map,nums[begin])
//				}
//				begin++
//			}
//		}
//		end++
//	}
//	return max_sum
//}