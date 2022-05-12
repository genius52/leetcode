package array

func plusOne(digits []int) []int {
	var l int = len(digits)
	var upgrade bool = true
	for i := l - 1;i >= 0;i--{
		val := digits[i]
		if upgrade{
			val++
		}
		digits[i] = val % 10
		if val >= 10{
			upgrade = true
		}else{
			upgrade = false
			break
		}
	}
	if upgrade{
		digits = append([]int{1},digits...)
	}
	return digits
}