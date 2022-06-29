package array

func MaximumXOR(nums []int) int {
	var res int = 0
	for i := 31;i >= 0;i--{
		var find bool = false
		for _,n := range nums{
			if (n | (1 << i)) == n{
				find = true
				break
			}
		}
		if find{
			res |= 1 << i
		}
	}
	return res
}