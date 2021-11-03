package string_issue

func CanConvertString(s string, t string, k int) bool {
	var l1 int = len(s)
	var l2 int = len(t)
	if l1 != l2{
		return false
	}
	var used [26]int
	for i := 0;i < l1;i++{
		diff := int(t[i]) - int(s[i])
		if diff == 0{
			continue
		}
		if diff < 0{
			diff += 26
		}
		move := diff + 26 * used[diff]
		if move > k{
			return false
		}
		used[diff]++
	}
	return true
}