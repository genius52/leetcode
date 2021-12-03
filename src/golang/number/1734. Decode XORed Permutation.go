package number

//Input: encoded = [3,1]
//Output: [1,2,3]
//Explanation: If perm = [1,2,3], then encoded = [1 XOR 2,2 XOR 3] = [3,1]
func Decode2(encoded []int) []int {
	var l int = len(encoded)
	var res []int = make([]int, l+1)
	var xor int = 0
	var total int = 0
	for i := 1; i <= l+1; i++ {
		total ^= i
	}
	for i := 0; i < l; i += 2 {
		xor = xor ^ encoded[i]
	}
	res[l] = total ^ xor
	for i := l - 1; i >= 0; i-- {
		res[i] = res[i+1] ^ encoded[i]
	}
	return res
}
