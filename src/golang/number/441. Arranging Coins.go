package number

func arrangeCoins(n int) int {
	var res int = 0
	var cnt int = 1
	for n > 0{
		n -= cnt
		if n < 0{
			break
		}
		cnt++
		res++
	}
	return res
}