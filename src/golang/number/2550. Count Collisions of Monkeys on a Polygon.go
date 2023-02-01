package number


func mypow(n int)int{
	var MOD int = 1e9 + 7
	var res int = 1
	var base int = 2
	for n > 0{
		if n | 1 == n{
			res = res * base % MOD
		}
		base = base * base % MOD
		n /= 2
	}
	return res
}

func monkeyMove(n int) int {
	var MOD int = 1e9 + 7
	return (mypow(n) - 2 + MOD) % MOD
}