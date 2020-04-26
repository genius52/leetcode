package array

import "math"

//Input: nums = [1,5,9,1,5,9], k = 2, t = 3
//Output: false
func ContainsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	l := len(nums)
	for i := 0;i < l;i++{
		for j := i + 1;j <= i + k;j++{
			if j >= l{
				break
			}
			if math.Abs(float64(nums[i]) - float64(nums[j])) <= float64(t){
				return true
			}
		}
	}
	return false
}
