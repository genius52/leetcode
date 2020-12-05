package string_issue

func CamelMatch(queries []string, pattern string) []bool {
	var pattern_len int = len(pattern)
	var l int = len(queries)
	var res []bool = make([]bool,l)
	for i := 0;i < l;i++{
		cur_s := queries[i]
		var cur_len int = len(cur_s)
		var p1 int = 0
		var p2 int = 0
		for p1 < cur_len && p2 < pattern_len{
			if cur_s[p1] != pattern[p2]{
				if pattern[p2] >= 'A' && pattern[p2] <= 'Z' &&
					cur_s[p1] >= 'A' && cur_s[p1] <= 'Z'{
					break
				}
				p1++
			}else{
				p2++
				p1++
			}
		}
		if p2 == pattern_len{
			var find_upper bool = false
			for p1 < cur_len{
				if cur_s[p1] >= 'A' && cur_s[p1] <= 'Z'{
					find_upper = true
					break
				}
				p1++
			}
			if find_upper{
				res[i] = false
			}else{
				res[i] = true
			}
		}else{
			res[i] = false
		}
	}
	return res
}