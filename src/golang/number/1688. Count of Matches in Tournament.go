package number

func numberOfMatches(n int) int {
	var res int = 0
	for n > 1{
		res += n / 2
		n = (n + 1) / 2
	}
	return res
}

//Lose game number
func numberOfMatches2(n int) int{
	return n - 1
}