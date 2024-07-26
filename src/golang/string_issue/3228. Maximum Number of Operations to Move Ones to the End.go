package string_issue

func maxOperations(s string) int {
	var l int = len(s)
	var res int = 0
	var one_cnt int = 0
	var moved bool = false
	for i := 0; i < l-1; i++ {
		if s[i] == '0' && !moved {
			res += one_cnt
			moved = true
		} else if s[i] == '1' {
			one_cnt++
			moved = false
		}
	}
	return res
}
