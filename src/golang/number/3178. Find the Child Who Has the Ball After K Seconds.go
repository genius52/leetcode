package number

func numberOfChild(n int, k int) int {
	round := k / (n - 1)
	m := k % (n - 1)
	if round|1 == round { //odd number
		return n - 1 - m
	} else {
		return m
	}
}
