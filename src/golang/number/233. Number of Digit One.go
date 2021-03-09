package number

//Input: n = 13
//Output: 6
func cacl_one(n int)int{
	var ones int = 0
	for n != 0{
		mod := n % 10
		if mod == 1{
			ones++
		}
		n = n / 10
	}
	return ones
}

func countDigitOne(n int) int {
	var total int = 0
	for i := 1;i <= n;i++{
		total += cacl_one(i)
	}
	return total
}