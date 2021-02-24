package number

func singleNumber(nums []int) int {
	var res int = 0
	for _,n := range nums{
		res ^= n
	}
	return res
}