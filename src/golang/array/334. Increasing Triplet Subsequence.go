package array

import "math"

//Return true if there exists i, j, k
//such that arr[i] < arr[j] < arr[k] given 0 ≤ i < j < k ≤ n-1 else return false.
//Note: Your algorithm should run in O(n) time complexity and O(1) space complexity.
func increasingTriplet(nums []int) bool {
	l := len(nums)
	if l <= 2{
		return false
	}
	var min int = nums[0]
	var mid int = math.MaxInt32
	for i := 1;i < len(nums);i++ {
		if nums[i] > mid{
			return true
		}
		if nums[i] > min{
			mid = nums[i]
		}else{
			min = nums[i]
		}
	}
	return false
}

func IncreasingTriplet2(nums []int) bool {
	var l int = len(nums)
	var dp_left_min []int = make([]int,l)
	var dp_right_max []int = make([]int,l)
	dp_left_min[0] = nums[0]
	dp_right_max[l - 1] = nums[l - 1]
	for i := 1;i < l;i++{
		dp_left_min[i] = min_int(dp_left_min[i - 1],nums[i])
		dp_right_max[l - 1 - i] = max_int(dp_right_max[l - i],nums[l - 1 - i])
	}
	for i := 1;i < l - 1;i++{
		if nums[i] > dp_left_min[i - 1] && nums[i] < dp_right_max[i + 1]{
			return true
		}
	}
	return false
}
