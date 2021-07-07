package string_issue

func IsLongPressedName(name string, typed string) bool {
	var l1 int = len(name)
	var l2 int = len(typed)
	if l2 < l1{
		return false
	}
	var left1 int = 0
	var left2 int = 0
	for left1 < l1 && left2 < l2{
		var right1 int = left1
		for right1 < l1 && name[right1] == name[left1]{
			right1++
		}
		var right2 int = left2
		for right2 < l2 && typed[right2] == typed[left2]{
			right2++
		}
		if name[left1] != typed[left2] || (right2 - left2) < (right1 - left1){
			return false
		}
		left1 = right1
		left2 = right2
	}
	if left1 == l1 && left2 == l2{
		return true
	}
	return false
}