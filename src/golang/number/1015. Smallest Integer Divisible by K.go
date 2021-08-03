package number

func SmallestRepunitDivByK(K int) int {
	if (K | 1) != K || (K % 5) == 0{
		return -1
	}
	n := 1
	for i := 1;i <= K;i++{
		mod := n % K
		if mod == 0{
			return i
		}
		n = (mod * 10 + 1) % K//只需保留余数，因为其他部分肯定能被K整除
	}
	return -1
}