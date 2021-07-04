package array

func smallestRangeI(nums []int, k int) int {
	var l int = len(nums)
	if l <= 1{
		return 0
	}
	var small int = 2147483647
	var big int = -2147483648
	for i := 0;i < l;i++{
		if nums[i] > big{
			big = nums[i]
		}
		if nums[i] < small{
			small = nums[i]
		}
	}
	diff := big - small
	if diff <= k * 2{
		return 0
	}
	return diff - k * 2
}