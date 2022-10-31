package number

//nums = [3,7,8,1,1,5], space = 2
func destroyTargets(nums []int, space int) int {
	var l int = len(nums)
	var record map[int]int = make(map[int]int)//mod space, count
	for i := 0;i < l;i++{
		record[nums[i] % space]++
	}
	//sort.Ints(nums)
	var max_cnt int = 0
	var max_cnt_num int = 0
	for i := 0;i < l;i++{
		key := nums[i] % space
		if record[key] > max_cnt{
			max_cnt = record[key]
			max_cnt_num = nums[i]
		}else if record[key] == max_cnt{
			if nums[i] < max_cnt_num{
				max_cnt_num = nums[i]
			}
		}
	}
	return max_cnt_num
}