package string_issue


func checkValidString(s string) bool {
	var left []int
	var star []int
	for i := 0;i < len(s);i++{
		if s[i] == '('{
			left = append(left, i)
		}else if s[i] == '*'{
			star = append(star,i)
		}else{
			if len(left) == 0 && len(star) == 0{
				return false
			}
			if len(left) > 0{
				left = left[:len(left) - 1]
			}else if len(star) > 0{
				star = star[:len(star) - 1]
			}
		}
	}
	for len(left) > 0 && len(star) > 0{
		l := left[len(left) - 1]
		s := star[len(star) - 1]
		left = left[:len(left) - 1]
		star = star[:len(star) - 1]
		if l > s{
			return false
		}
	}
	if len(left) > 0{
		return false
	}
	return true
}