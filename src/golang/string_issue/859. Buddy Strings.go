package string_issue

func buddyStrings(s string, goal string) bool {
	var l1 int = len(s)
	var l2 int = len(goal)
	if l1 != l2{
		return false
	}
	if s == goal{
		var cnt [26]int
		for i := 0;i < l1;i++{
			cnt[s[i] - 'a']++
			if cnt[s[i] - 'a'] > 1{
				return true
			}
		}
		return false
	}
	var first_index int = -1
	var second_index int = -1
	for i := 0;i < l1;i++{
		if s[i] == goal[i]{
			continue
		}
		if first_index == -1{
			first_index = i
		}else{
			if second_index == -1{
				second_index = i
			}else{
				return false
			}
		}
	}
	if first_index == -1 || second_index == -1{
		return false
	}
	if s[first_index] == goal[second_index] && s[second_index] == goal[first_index]{
		return true
	}
	return false
}