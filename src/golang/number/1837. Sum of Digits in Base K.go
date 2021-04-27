package number

//Input: n = 34, k = 6
//Output: 9
//Explanation: 34 (base 10) expressed in base 6 is 54. 5 + 4 = 9.
func sumBase(n int, k int) int {
	var res int = 0
	for n > 0{
		res += (n % k)
		n = n / k
	}
	return res
}