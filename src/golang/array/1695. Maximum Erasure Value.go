package array

//Input: nums = [4,2,4,5,6]
//Output: 17
//Explanation: The optimal subarray here is [2,4,5,6].
func MaximumUniqueSubarray(nums []int) int {
	var l int = len(nums)
	var begin int = 0
	var end int = 0
	var record [10001]int
	var dup_map map[int]bool = make(map[int]bool)
	var cur_sum int = 0
	var max_sum int = 0
	for end < l{
		cur_sum += nums[end]
		record[nums[end]]++
		if record[nums[end]] > 1{
			dup_map[nums[end]] = true
		}
		if len(dup_map) == 0{
			if cur_sum > max_sum{
				max_sum = cur_sum
			}
		}else{
			for len(dup_map) > 0 && begin < end{
				cur_sum -= nums[begin]
				record[nums[begin]]--
				if record[nums[begin]] == 1{
					delete(dup_map,nums[begin])
				}
				begin++
			}
		}
		end++
	}
	return max_sum
}