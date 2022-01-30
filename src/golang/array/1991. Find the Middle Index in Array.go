package array

func findMiddleIndex(nums []int) int {
	var l int = len(nums)
	var sum int = 0
	var prefix_sum []int = make([]int,l + 1)
	for i := 0;i < l;i++{
		prefix_sum[i + 1] = prefix_sum[i] + nums[i]
		sum += nums[i]
	}
	for i := 0;i < l;i++{
		if prefix_sum[i] == (sum - prefix_sum[i] - nums[i]){
			return i
		}
	}
	return -1
}

//O(1) space solution
func findMiddleIndex2(nums []int) int{
	var l int = len(nums)
	var sum int = 0
	var left_sum int = 0
	for i := 0;i < l;i++{
		sum += nums[i]
	}
	for i := 0;i < l;i++{
		sum -= nums[i]
		if left_sum == sum{
			return i
		}
		left_sum += nums[i]
	}
	return -1
}