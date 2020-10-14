package string_issue

//Input: start = "RXXLRXRXL", end = "XRLXXRRLX"
//Output: true
//replacing one occurrence of "XL" with "LX", or replacing one occurrence of "RX" with "XR"
func CanTransform(start string, end string) bool {
	var l int = len(start)
	var p1 int = 0
	var p2 int = 0
	for p1 < l || p2 < l{
		for p1 < l && start[p1] == 'X'{
			p1++
		}
		for p2 < l && end[p2] == 'X'{
			p2++
		}
		if p1 < l && p2 < l && start[p1] != end[p2]{
			return false
		}
		if p1 < l && p2 < l && start[p1] == end[p2]{
			if start[p1] == 'L' && p1 < p2 {
				return false
			}
			if start[p1] == 'R' && p1 > p2{
				return false
			}
		}
		if p1 == l && p2 == l{
			return true
		}
		if p1 == l || p2 == l{
			return false
		}
		p1++
		p2++
	}
	if p1 == l && p2 == l{
		return true
	}
	return false
}