package string_issue

//"bbbacddceeb"
//"ceebbbbacdd"
func RotateString(s string, goal string) bool {
	var l1 int = len(s)
	var l2 int = len(goal)
	if l1 != l2{
		return false
	}
	if l1 == 0 && l2 == 0{
		return true
	}
	for j := 0;j < l2;j++{
		if goal[j] != s[0]{
			continue
		}
		var visit1 int = 0
		var visit2 int = j
		for visit1 < l1{
			if goal[visit2] != s[visit1]{
				break
			}
			visit1++
			visit2 = (visit2 + 1) % l2
		}
		if visit1 == l1{
			return true
		}
	}
	return false
}