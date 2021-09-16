package array

//Input: nums = [1,2,3,4]
//Output: [2,4,4,4]
//Explanation: The first pair [1,2] means we have freq = 1 and val = 2 so we generate the array [2].
//The second pair [3,4] means we have freq = 3 and val = 4 so we generate [4,4,4].
func decompressRLElist(nums []int) []int {
	l := len(nums)
	var total_length int = 0
	for i := 0;i < l;i += 2{
		total_length += nums[i]
	}
	var res []int = make([]int,total_length)
	var cur_pos int = 0
	for i := 0;i < l;i += 2{
		cnt := nums[i]
		val := nums[i + 1]
		for j := 0;j < cnt;j++{
			res[cur_pos] = val
			cur_pos++
		}
	}
	return res
}