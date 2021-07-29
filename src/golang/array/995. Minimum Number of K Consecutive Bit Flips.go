package array

import "container/list"

func MinKBitFlips(nums []int, k int) int{
	var l int = len(nums)
	var q list.List
	var steps int = 0
	for i := 0;i < l;i++{
		flip := q.Len() % 2
		cur := nums[i] ^ flip
		for q.Len() > 0{
			top := q.Front().Value.(int)
			if top <= i{
				q.Remove(q.Front())
			}else{
				break
			}
		}
		if cur == 1{
			continue
		}else{
			if k > 1{
				q.PushBack(i + k - 1)
			}
		}
		steps++
	}
	if q.Len() > 0{
		return -1
	}
	return  steps
}

//TLE
//func MinKBitFlips(nums []int, k int) int {
//	var l int = len(nums)
//	var steps int = 0
//	for i := 0;(i + k) <= l;i++{
//		if nums[i] == 1{
//			continue
//		}
//		for j := i;j < i + k;j++{
//			if nums[j] == 1{
//				nums[j] = 0
//			}else{
//				nums[j] = 1
//			}
//		}
//		steps++
//	}
//	for i := l - 1;i >= 0;i--{
//		if nums[i] != 1{
//			return -1
//		}
//	}
//	return steps
//}