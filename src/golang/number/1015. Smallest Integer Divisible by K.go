package number

func SmallestRepunitDivByK(K int) int {
	if (K | 1) != K || (K % 5) == 0{
		return -1
	}
	n := 0
	for i := 1;i <= K;i++{
		n = (n * 10 + 1) % K
		if n == 0{
			return i
		}
	}
	return -1
}