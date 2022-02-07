package array

func SortEvenOdd(nums []int) []int {
	var l int = len(nums)
	for i := 0;i < l;i += 2{
		for j := 0;j + 2 < l - i;j += 2{
			if nums[j] > nums[j + 2]{
				nums[j],nums[j + 2] = nums[j + 2],nums[j]
			}
		}
	}
	for i := 0;i < l;i += 2{
		for j := 1;j + 2 < l - i;j += 2{
			if nums[j] < nums[j + 2]{
				nums[j],nums[j + 2] = nums[j + 2],nums[j]
			}
		}
	}
	return nums
}