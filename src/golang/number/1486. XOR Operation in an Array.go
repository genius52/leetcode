package number

//nums[i] = start + 2*i
func xorOperation(n int, start int) int {
	var res int = 0
	for i := 0;i < n;i++{
		res = res ^ (start + 2 * i)
	}
	return res
}