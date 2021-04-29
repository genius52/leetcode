package array

//Input: [1,12,-5,-6,50,3], k = 4
//Output: 12.75
//Explanation: Maximum average is (12-5-6+50)/4 = 51/4 = 12.75
func FindMaxAverage(nums []int, k int) float64 {
	var l int = len(nums)
	if l < k{
		return 0
	}
	var sum int = 0
	for i := 0;i < k;i++{
		sum += nums[i]
	}
	var res int = sum
	for i := k;i < l;i++{
		sum -= nums[i - k]
		sum += nums[i]
		if sum > res{
			res = sum
		}
	}
	return float64(res)/float64(k)
}