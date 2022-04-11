package array

//Input:nums = [1,1,1], k = 2
//Output: 2
func SubarraySum(nums []int, k int) int {
	var res int = 0
	var record map[int]int = make(map[int]int)
	var sum int = 0
	for _,n := range nums{
		sum += n
		if sum == k{
			res += 1
		}
		if _,ok := record[sum - k];ok{
			res += record[sum - k]
		}
		record[sum]++
	}
	return res
}