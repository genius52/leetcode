package number

import (
	"math"
)

//bucket sort
func smallestDistancePair(nums []int, k int) int{
	var data [1000000]int
	var l int = len(nums)
	for i := 0;i < l;i++{
		for j := i + 1;j < l;j++{
			data[int(math.Abs(float64(nums[i] - nums[j])))]++
		}
	}
	for i := 0;i < 1000000;i++{
		if data[i] == 0{
			continue
		}
		if data[i] < k{
			k -= data[i]
		}else{
			return i
		}
	}
	return -1
}

//func count_smaller(nums []int,diff int)int{
//	return 0
//}
//
//func smallestDistancePair(nums []int, k int) int {
//	var l int = len(nums)
//	sort.Ints(nums)
//	var left int = 2147483647
//	for i := 0;i < l;i++{
//		var diff int = int(math.Abs(float64(nums[i] - nums[i - 1])))
//		if diff < left{
//			left = diff
//		}
//	}
//	var right int = nums[l - 1] - nums[0]
//	for left < right{
//		var mid int = left + (right - left)/2
//		var res int = count_smaller(nums,mid)
//		if res == k{
//			return mid
//		}else if res < k{
//			left = mid + 1
//		}else{
//			right = mid - 1
//		}
//	}
//	return left
//}