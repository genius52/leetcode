package number

func numberOfCuts(n int) int {
	if n == 1{
		return 0
	}
	if n | 1 == n{ //odd number
		return n
	}else{
		return n / 2
	}
}