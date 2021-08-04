package number

//Input: n = 3
//Output: "111"
//Explantion: (-2) ^ 2 + (-2) ^ 1 + (-2) ^ 0 = 3
func BaseNeg2(N int) string {
	if N == 0{
		return "0"
	}
	var res string
	for N != 0{
		if (N % -2) == 0{
			res = "0" + res
		}else{
			res = "1" + res
			N--
		}
		N = N / -2
	}
	return res
}