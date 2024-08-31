package array

func ResultsArray2(nums []int, k int) []int {
	var l int = len(nums)
	var res []int = make([]int, l-k+1)
	var increase_cnt int = 0
	for i := 0; i < k-1; i++ {
		if nums[i]+1 == nums[i+1] {
			increase_cnt++
		}
	}
	if increase_cnt == k-1 {
		res[0] = nums[k-1]
	} else {
		res[0] = -1
	}
	for i := 1; i+k <= l; i++ {
		if nums[i-1] == nums[i]-1 {
			increase_cnt--
		}
		if nums[i+k-2] == nums[i+k-1]-1 {
			increase_cnt++
		}
		if increase_cnt == k-1 {
			res[i] = nums[i+k-1]
		} else {
			res[i] = -1
		}
	}
	return res
}
