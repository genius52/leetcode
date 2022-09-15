package number

func calc_superPow(n int,times int)int{
	MOD := 1337
	var res int = 1
	for i := 0;i < times;i++{
		res = res * n % MOD
	}
	return res
}

//a = 1, b = [4,3,3,8,5,2]
func superPow(a int, b []int) int {
	if a == 1{
		return 1
	}
	var l int = len(b)
	if l == 0{
		return 1
	}
	last := b[l - 1]
	b = b[:l - 1]
	MOD := 1337
	return calc_superPow(superPow(a,b),10) * calc_superPow(a,last) % MOD
}