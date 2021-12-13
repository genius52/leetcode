package string_issue

func checkOnesSegment(s string) bool {
	var l int = len(s)
	var segment int = 0
	var last_is_one bool = false
	for i := 0; i < l; i++ {
		if s[i] == '1' {
			if !last_is_one {
				segment++
			}
			if segment > 1 {
				return false
			}
			last_is_one = true
		} else {
			last_is_one = false
		}
	}
	return true
}
