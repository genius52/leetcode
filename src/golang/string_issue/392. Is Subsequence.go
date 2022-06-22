package string_issue

//func dp_isSubsequence(s string,l1 int,pos1 int,t string,l2 int,pos2 int)bool{
//	if pos2 >= l2 && l1
//}

func isSubsequence(s string, t string) bool {
	var l1 int = len(s)
	var l2 int = len(t)
	if l1 > l2{
		return false
	}
	var pos1 int = 0
	var pos2 int = 0
	for pos1 < l1 && pos2 < l2{
		if s[pos1] == t[pos2]{
			pos1++
			pos2++
		}else{
			pos2++
		}
	}
	return pos1 == l1
}