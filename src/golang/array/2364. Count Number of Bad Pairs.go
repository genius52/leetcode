package array

func countBadPairs(nums []int) int64 {
	var l int = len(nums)
	var res int64 = 0
	var record map[int]int = make(map[int]int)
	for i := 0;i < l;i++{
		diff := nums[i] - i
		var cnt int = i
		if val,ok := record[diff];ok{
			cnt -= val
		}
		res += int64(cnt)
		record[nums[i] - i]++
	}
	return res
}