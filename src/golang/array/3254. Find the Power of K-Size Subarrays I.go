package array

func ResultsArray(nums []int, k int) []int {
	var l int = len(nums)
	var res []int = make([]int, l-k+1)
	for i := 0; i < l-k+1; i++ {
		res[i] = -1
	}
	for i := 0; i+k <= l; i++ {
		var increase bool = true
		for j := 0; j < k-1; j++ {
			if nums[i+j] != nums[i+j+1]-1 {
				increase = false
				break
			}
		}
		if increase {
			res[i] = nums[i+k-1]
		}
	}
	return res
}
