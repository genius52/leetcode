package number

import "strconv"

func countSymmetricIntegers(low int, high int) int {
	var res int = 0
	check := func(x int) bool {
		s := strconv.Itoa(x)
		var l int = len(s)
		if l%2 == 1 {
			return false
		}
		pre := 0
		suffix := 0
		for i := 0; i < l/2; i++ {
			pre += int(s[i] - '0')
			suffix += int(s[l-1-i] - '0')
		}
		return pre == suffix
	}
	for i := low; i <= high; i++ {
		if check(i) {
			res++
		}
	}
	return res
}
