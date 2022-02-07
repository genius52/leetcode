package array

func pivotArray(nums []int, pivot int) []int {
	var l int = len(nums)
	var res []int = make([]int,l)
	var idx int = 0
	for i := 0;i < l;i++{
		if nums[i] < pivot{
			res[idx] = nums[i]
			idx++
			}
	}
	for i := 0;i < l;i++{
		if nums[i] == pivot{
			res[idx] = nums[i]
			idx++
		}
	}
	for i := 0;i < l;i++{
		if nums[i] > pivot{
			res[idx] = nums[i]
			idx++
		}
	}
	return res
}