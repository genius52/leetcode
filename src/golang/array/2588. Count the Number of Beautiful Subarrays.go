package array

func BeautifulSubarrays(nums []int) int64 {
	var l int = len(nums)
	var res int64 = 0
	var record map[int]int64 = make(map[int]int64)
	record[0] = 1
	var cur int = 0
	for i := 0; i < l; i++ {
		cur ^= nums[i]
		if cnt, ok := record[cur]; ok {
			res += cnt
		}
		record[cur]++
	}
	return res
}
