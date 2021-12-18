package number

func arraySign(nums []int) int {
	var product int = 1
	for _, n := range nums {
		if n == 0 {
			return 0
		}
		if n < 0 {
			product = -product
		}
	}
	return product
}
