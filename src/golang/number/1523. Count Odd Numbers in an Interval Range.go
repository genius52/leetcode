package number

func CountOdds(low int, high int) int {
	var res int = (high - low) / 2
	if low % 2 == 1 || high % 2 == 1{
		res++
	}
	return res
}