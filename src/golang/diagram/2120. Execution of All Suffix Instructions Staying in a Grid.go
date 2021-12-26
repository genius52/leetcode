package diagram

func executeInstructions(n int, startPos []int, s string) []int {
	var l int = len(s)
	var res []int = make([]int, l)
	for i := l - 1; i >= 0; i-- {
		r := startPos[0]
		c := startPos[1]
		steps := 0
		for j := i; j < l; j++ {
			if s[j] == 'L' {
				c--
			} else if s[j] == 'R' {
				c++
			} else if s[j] == 'U' {
				r--
			} else if s[j] == 'D' {
				r++
			}
			if r < 0 || r >= n || c < 0 || c >= n {
				break
			}
			steps++
		}
		res[i] = steps
	}
	return res
}
