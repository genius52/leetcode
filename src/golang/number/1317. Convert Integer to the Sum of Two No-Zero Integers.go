package number

func check_zero(num int)bool{
	for num != 0{
		if (num % 10) == 0{
			return false
		}
		num = num / 10
	}
	return true
}

func getNoZeroIntegers(n int) []int {
	low := 1
	high := n - 1

	for low <= high{
		if check_zero(low) && check_zero(high){
			return []int{low,high}
		}
		low++
		high--
	}
	return []int{0,0}
}