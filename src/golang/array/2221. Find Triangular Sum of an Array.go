package array

func triangularSum(nums []int) int {
	cur := nums
	for len(cur) > 1{
		var cur_len int = len(cur)
		var next []int = make([]int,cur_len - 1)
		for i := 0;i < cur_len - 1;i++{
			next[i] = (cur[i] + cur[i + 1]) % 10
		}
		cur = next
	}
	return cur[0]
}