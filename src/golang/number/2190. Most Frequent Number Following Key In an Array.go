package number

func mostFrequent(nums []int, key int) int {
	var record map[int]int = make(map[int]int)
	var max_cnt int = 0
	var res int = 0
	for i := 1;i < len(nums);i++{
		if nums[i - 1] == key{
			record[nums[i]]++
			if record[nums[i]] > max_cnt{
				max_cnt = record[nums[i]]
				res = nums[i]
			}
		}
	}
	return res
}