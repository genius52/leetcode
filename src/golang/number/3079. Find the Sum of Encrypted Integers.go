package number

func sumOfEncryptedInt(nums []int) int {
	var sum int = 0
	for _, n := range nums {
		var max_digit int = 0
		var offset int = 0
		for n > 0 {
			m := n % 10
			if m > max_digit {
				max_digit = m
			}
			n /= 10
			offset++
		}
		for offset > 0 {
			sum += max_digit
			max_digit *= 10
			offset--
		}
	}
	return sum
}
