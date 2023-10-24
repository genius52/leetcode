package number

// Define two integers, num1 and num2, as follows:
//
// num1: The sum of all integers in the range [1, n] that are not divisible by m.
// num2: The sum of all integers in the range [1, n] that are divisible by m.
// Return the integer num1 - num2.
func differenceOfSums(n int, m int) int {
	var num1 int = 0
	var num2 int = 0
	for i := 1; i <= n; i++ {
		if i%m != 0 {
			num1 += i
		} else {
			num2 += i
		}
	}
	return num1 - num2
}
