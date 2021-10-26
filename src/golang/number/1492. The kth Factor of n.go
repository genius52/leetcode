package number

func kthFactor(n int, k int) int {
	if k > n{
		return -1
	}
	for i := 1;i <= n / 2;i++{
		if n % i == 0{
			k--
			if k == 0{
				return i
			}
		}
	}
	if k == 1{
		return n
	}else{
		return -1
	}
}