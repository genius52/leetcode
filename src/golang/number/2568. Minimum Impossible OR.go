package number

import "sort"

func MinImpossibleOR(nums []int) int {
	var l int = len(nums)
	sort.Ints(nums)
	var offset int = 0
	var i int = 0
	var target int = 1
	for offset < 32 && i < l {
		if nums[i] > target {
			return target
		} else if nums[i] == target {
			target <<= 1
			offset++
		} else {
			i++
		}
	}
	return target
}

//func MinImpossibleOR(nums []int) int {
//	var l int = len(nums)
//	sort.Ints(nums)
//	for i := 0; i < 32; i++ {
//		target := 1 << i
//		var mask []bool = make([]bool, i+1)
//		for j := 0; j < l; j++ {
//			if nums[j] > target {
//				break
//			}
//			var n int = nums[j]
//			var offset int = 0
//			for n > 0 {
//				if n|1 == n {
//					mask[offset] = true
//				}
//				n >>= 1
//				offset++
//			}
//		}
//		var find bool = true
//		for k := 0; k <= i; k++ {
//			if !mask[k] {
//				find = false
//				break
//			}
//		}
//		if !find {
//			return target
//		}
//	}
//	return 0
//}
