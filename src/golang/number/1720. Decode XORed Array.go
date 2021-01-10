package number

//Input: encoded = [1,2,3], first = 1
//Output: [1,0,2,1]
func decode(encoded []int, first int) []int {
	var l int = len(encoded)
	var res []int = make([]int,l + 1)
	res[0] = first
	for i := 1;i <= l;i++{
		res[i] = encoded[i - 1] ^ res[i - 1]
	}
	return res
}