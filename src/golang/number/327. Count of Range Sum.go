package number

//Input: nums = [-2,5,-1], lower = -2, upper = 2
//Output: 3
//Explanation: The three ranges are: [0,0], [2,2], and [0,2] and their respective sums are: -2, -1, 2.
func CountRangeSum(nums []int, lower int, upper int) int {
	var l int = len(nums)
	var prefix []int64 = make([]int64,l + 1)
	prefix[0] = 0
	for i := 1;i <= l;i++{
		prefix[i] = int64(nums[i - 1]) + prefix[i - 1]
	}
	var count int = 0
	for i := 0;i <= l;i++{
		for j := i + 1;j <= l;j++{
			var sum int64 = prefix[j] - prefix[i]
			if sum >= int64(lower) && sum <= int64(upper){
				count++
			}
		}
	}
	return count
}