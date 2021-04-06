package array

//Input: nums = [3,1,4,2]
//Output: true
//Explanation: There is a 132 pattern in the sequence: [1, 4, 2].
func find132pattern(nums []int) bool {
	var l int = len(nums)
	if l < 3{
		return false
	}
	var dp_min []int = make([]int,l)
	dp_min[0] = nums[0]
	//var min_num int = 2147483647
	for i := 1;i < l;i++{
		dp_min[i] = min_int(nums[i],dp_min[i - 1])
		if nums[i] == dp_min[i]{
			continue
		}
		for j := i - 1;j > 0;j--{
			if dp_min[j - 1] >= nums[i]{
				break
			}
			if nums[j] > nums[i] && nums[j] > dp_min[j - 1] && dp_min[j - 1] < nums[i]{
				return true
			}
		}
	}
	return false
}