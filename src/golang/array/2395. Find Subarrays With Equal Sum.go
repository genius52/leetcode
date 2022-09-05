package array

func findSubarrays(nums []int) bool {
	var l int = len(nums)
	var record map[int]bool = make(map[int]bool)
	for i := 1;i < l;i++{
		sum := nums[i] + nums[i - 1]
		if _,ok := record[sum];ok{
			return true
		}
		record[sum] = true
	}
	return false
}