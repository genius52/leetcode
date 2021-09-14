package number

func sumZero(n int) []int {
	var res []int = make([]int,n)
	var left int = 0
	var right int = n - 1
	var i int = 1
	for left < right{
		res[left] = i
		res[right] = -i
		left++
		right--
		i++
	}
	return res
}