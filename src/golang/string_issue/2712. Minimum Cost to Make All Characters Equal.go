package string_issue

func min_int64(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func MinimumCost(s string) int64 {
	var l int = len(s)
	var prefix [][2]int64 = make([][2]int64, l) //prefix[i][0] 把0到i 都转成0的代价
	var suffix [][2]int64 = make([][2]int64, l)
	for i := 0; i < l; i++ {
		if s[i] == '0' {
			if i != 0 {
				prefix[i][1] = int64(i+1) + prefix[i-1][0]
				prefix[i][0] = prefix[i-1][0]
			} else {
				prefix[i][1] = 1
			}
		} else if s[i] == '1' {
			if i != 0 {
				prefix[i][0] = int64(i+1) + prefix[i-1][1]
				prefix[i][1] = prefix[i-1][1]
			} else {
				prefix[i][0] = 1
			}
		}
	}
	for i := l - 1; i >= 0; i-- {
		if s[i] == '0' {
			if i == (l - 1) {
				suffix[i][1] = 1
			} else {
				suffix[i][1] = int64(l-i) + suffix[i+1][0]
				suffix[i][0] = suffix[i+1][0]
			}
		} else if s[i] == '1' {
			if i == (l - 1) {
				suffix[i][0] = 1
			} else {
				suffix[i][0] = int64(l-i) + suffix[i+1][1]
				suffix[i][1] = suffix[i+1][1]
			}
		}
	}
	var res int64 = 1<<63 - 1
	for i := 0; i < l; i++ {
		if i == 0 {
			val := min_int64(suffix[i][0], suffix[i][1])
			res = min_int64(res, val)
		} else if i == (l - 1) {
			val := min_int64(prefix[i][0], prefix[i][1])
			res = min_int64(res, val)
		} else {
			val := min_int64(prefix[i-1][0]+suffix[i][0], prefix[i-1][1]+suffix[i][1])
			res = min_int64(res, val)
		}
	}
	return res
}
