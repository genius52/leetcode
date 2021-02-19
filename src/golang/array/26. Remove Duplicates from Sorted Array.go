package array

//Input: nums = [0,0,1,1,1,2,2,3,3,4]
//Output: 5, nums = [0,1,2,3,4]
func RemoveDuplicates(nums []int) int {
	var l int = len(nums)
	var dup_cnt int = 0
	for i := 0;i < l;i++{
		if nums[i] == nums[i - 1]{
			dup_cnt++
		}else{
			nums[i - dup_cnt] = nums[i]
		}
	}
	return l - dup_cnt
}