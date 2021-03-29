package string_issue

func evaluate(s string, knowledge [][]string) string {
	var l int = len(s)
	var dict map[string]string = make(map[string]string)
	for _,know := range knowledge{
		dict[know[0]] = know[1]
	}
	var left int = 0
	var res string
	for left < l {
		if s[left] != '('{
			res += string(s[left])
			left++
		}else{
			right := left
			for right < l{
				if s[right] != ')'{
					right++
				}else{
					break;
				}
			}
			sub := s[left+1:right]
			if _,ok := dict[sub];ok{
				res += dict[sub]
			}else{
				res += "?"
			}
			left = right + 1
		}
	}
	return res
}