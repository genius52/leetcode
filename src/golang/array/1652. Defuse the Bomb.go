package array

func Decrypt(code []int, k int) []int {
	var l int = len(code)
	var res []int = make([]int,l)
	if k == 0{
		return res
	}
	var dir int = 1
	if k < 0{
		dir = -1
	}
	for i := 0;i < l;i++{
		var total int = 0
		for j := 0;j != k;j += dir{
			total += code[(i + dir + j + l) % l]
		}
		res[i] = total
	}
	return res
}