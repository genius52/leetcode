package number

func hammingDistance(x int, y int) int {
	var res int = 0
	for i := 0;i < 31;i++{
		mask := 1 << i
		x1 := x & mask
		y1 := y & mask
		if x1 != y1{
			res++
		}
		if mask > x && mask > y{
			break
		}
	}
	return res
}