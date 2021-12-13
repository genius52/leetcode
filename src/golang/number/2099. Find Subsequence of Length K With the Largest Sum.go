package number

import "sort"

//Input: nums = [2,1,3,3], k = 2
//Output: [3,3]
//Explanation:
//The subsequence has the largest sum of 3 + 3 = 6.
func maxSubsequence(nums []int, k int) []int {
	var l int = len(nums)
	var record [][2]int = make([][2]int, l)
	for i := 0; i < l; i++ {
		record[i][0] = nums[i]
		record[i][1] = i
	}
	sort.Slice(record, func(i, j int) bool {
		return record[i][0] > record[j][0]
	})
	record = record[:k]
	sort.Slice(record, func(i, j int) bool {
		return record[i][1] < record[j][1]
	})
	var res []int = make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = record[i][0]
	}
	return res
}
