package array

func smallestEqual(nums []int) int {
	var l int = len(nums)
	for i := 0;i < l;i++{
		if nums[i] >= 10{
			continue
		}
		if i % 10 == nums[i]{
			return i
		}
	}
	return -1
}