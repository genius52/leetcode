package string_issue

func canChange(start string, target string) bool {
	var l int = len(start)
	var i int = 0
	var j int = 0
	for i < l || j < l{
		for i < l && start[i] == '_'{
			i++
		}
		for j < l && target[j] == '_'{
			j++
		}
		if i < l && j < l{
			if start[i] != target[j]{
				return false
			}
			if start[i] == 'L'{
				if i < j{
					return false
				}
			}else if start[i] == 'R'{
				if i > j{
					return false
				}
			}
			i++
			j++
		}else{
			break
		}
	}
	//return i == l && j == l
	return true
}