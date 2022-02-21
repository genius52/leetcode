package number

func sumOfThree(num int64) []int64 {
	var res []int64
	if num%3 != 0 {
		return res
	}
	res = []int64{num/3 - 1, num / 3, num/3 + 1}
	return res
}