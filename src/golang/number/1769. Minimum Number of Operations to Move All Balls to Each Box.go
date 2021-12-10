package number

func MinOperations2(boxes string) []int {
	var l int = len(boxes)
	var res []int = make([]int, l)
	var total int = 0
	for i := 0; i < l; i++ {
		if boxes[i] == '1' {
			total++
			res[0] += i
		}
	}
	var left int = 0
	if boxes[0] == '1' {
		left++
	}
	for i := 1; i < l; i++ {
		right := total - left
		res[i] = res[i-1] + left
		res[i] -= right
		if boxes[i] == '1' {
			left++
		}
	}
	return res
}

func minOperations(boxes string) []int {
	var l int = len(boxes)
	var res []int = make([]int, l)
	for i := 0; i < l; i++ {
		if boxes[i] == '0' {
			continue
		}
		for j := 0; j < l; j++ {
			res[j] += abs_int(i - j)
		}
	}
	return res
}
