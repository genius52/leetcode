package array

import "strconv"

func findTheArrayConcVal(nums []int) int64 {
	var l int = len(nums)
	var res int64 = 0
	var left int = 0
	var right int = l - 1
	for left < right{
		val,_ := strconv.Atoi(strconv.Itoa(nums[left]) + strconv.Itoa(nums[right]))
		res += int64(val)
		left++
		right--
	}
	if left == right{
		res += int64(nums[left])
	}
	return res
}