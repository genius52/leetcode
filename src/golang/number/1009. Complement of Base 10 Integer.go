package number

//Input: n = 10
//Output: 5
//Explanation: 10 is "1010" in binary, with complement "0101" in binary, which is 5 in base-10.
func bitwiseComplement(n int) int {
	mask := 1
	for mask < n{
		mask <<= 1
		mask += 1
	}
	return mask ^ n
}