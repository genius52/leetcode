package number

func differenceOfSum(nums []int) int {
	var sum int = 0
	var digit int = 0
	for _,n := range nums{
		sum += n
		for n > 0{
			digit += n % 10
			n /= 10
		}
	}
	res := sum - digit
	if res < 0{
		return -res
	}
	return res
}