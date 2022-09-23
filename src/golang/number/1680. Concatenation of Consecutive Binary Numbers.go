package number

//func recursive_concatenatedBinary(n int)int{
//	var MOD int = 1e9 + 7
//}

//f(0) = 0
//f(1) = 1     = (f(0) << 1) + 1  // len(1) = 1
//f(2) = 110   = (f(1) << 2) + 2  // len(2) = 2
//f(3) = 11011 = (f(2) << 2) + 3  // len(3) = 2
func ConcatenatedBinary(n int) int {
	var res int = 0
	var MOD int = 1e9 + 7
	var offset int = 1
	for i := 1;i <= n;i++{
		res = res << offset
		if (i & (i + 1)) == 0{ //1,2,4
			offset++
		}
		res = res | i
		res %= MOD
	}
	return int(res)
}