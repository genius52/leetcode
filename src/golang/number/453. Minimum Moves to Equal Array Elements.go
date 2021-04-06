package number

func minMoves(nums []int) int {
	var min_num int = 2147483647
	for _,n := range nums{
		if n < min_num{
			min_num = n
		}
	}
	var res int = 0
	for _,n := range nums{
		res += n - min_num
	}
	return res
}