package array

func MaxNonOverlapping(nums []int, target int) int {
	var l int = len(nums)
	var record map[int]bool = make(map[int]bool)
	var res int = 0
	var cur_sum int = 0
	for i := 0;i < l;i++{
		cur_sum += nums[i]
		if cur_sum == target{
			res++
			cur_sum = 0
			record = make(map[int]bool)
		}else if _,ok := record[cur_sum - target];ok{
			res++
			cur_sum = 0
			record = make(map[int]bool)
		}else{
			record[cur_sum] = true
		}
	}
	return res
}