package array

//Can be optimized by sliding window
func getAverages(nums []int, k int) []int {
	if k == 0{
		return nums
	}
	var l int = len(nums)
	var prefix []int = make([]int,l + 1)
	for i := 0;i < l;i++{
		prefix[i + 1] = prefix[i] + nums[i]
	}
	var res []int = make([]int,l)
	for i := 0;i < min_int(k,l);i++{
		res[i] = -1
	}
	for i := l - 1;i > max_int(l - 1 - k,0);i--{
		res[i] = -1
	}
	for i := k;i <= l - 1 - k;i++{
		sum := prefix[i + k + 1] - prefix[i - k]
		res[i] = sum / (k * 2 + 1)
	}
	return res
}