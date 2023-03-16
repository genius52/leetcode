package number

func passThePillow(n int, time int) int {
	mod := time % ((n - 1) * 2)
	if mod < n {
		return 1 + mod
	} else {
		mod %= n - 1
		return n - mod
	}
}
