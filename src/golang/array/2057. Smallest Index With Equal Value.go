package array

func smallestEqual(nums []int) int {
	var l int = len(nums)
	var res int = -1
	for i := 0;i < l;i++{
		if i % 10 == nums[i]{
			res = i
			break
		}
	}
	return res
}