package array

func check(nums []int) bool {
	var l int = len(nums)
	var cnt int = 0
	for i := 1; i < l; i++ {
		if nums[i] <= nums[i-1] {
			cnt++
		}
	}
	if cnt > 1 {
		return false
	}
	if cnt == 1 {
		return nums[l-1] <= nums[0]
	}
	return true
}

//func Check(nums []int) bool {
//	var l int = len(nums)
//	if l <= 2 {
//		return true
//	}
//	var decrease_index int = -1
//	for i := 1; i < l; i++ {
//		if nums[i] < nums[i-1] {
//			decrease_index = i
//			break
//		}
//	}
//	if decrease_index == -1 {
//		return true
//	}
//	var cnt int = 0
//	var begin int = (decrease_index + 1) % l
//	for cnt < l-1 {
//		if nums[begin] < nums[(begin-1+l)%l] {
//			return false
//		}
//		begin = (begin + 1) % l
//		cnt++
//	}
//	return true
//}
