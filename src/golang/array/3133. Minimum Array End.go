package array

// n: 数组长度
// x: 位运算 &  结果
func MinEnd(n int, x int) int64 {
	var res int64 = int64(x)
	for n > 1 {
		res = (res + 1) | int64(x)
		n--
	}
	return res
}
