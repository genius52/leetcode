package string_issue

func areAlmostEqual2(s1 string, s2 string) bool {
	var l1 int = len(s1)
	var l2 int = len(s2)
	if l1 != l2 {
		return false
	}
	var wrong_cnt int = 0
	var a1, a2, b1, b2 uint8 = 0, 0, 0, 0
	for i := 0; i < l1; i++ {
		if s1[i] != s2[i] {
			wrong_cnt++
			if wrong_cnt > 2 {
				return false
			}
			if a1 == 0 {
				a1 = s1[i]
				b1 = s2[i]
			} else {
				a2 = s1[i]
				b2 = s2[i]
				if a1 != b2 || a2 != b1 {
					return false
				}
			}
		}
	}
	return a1 == b2 && a2 == b1
}

func areAlmostEqual(s1 string, s2 string) bool {
	var l1 int = len(s1)
	var l2 int = len(s2)
	if l1 != l2 {
		return false
	}
	var wrong_pos int = 0
	var cnt1, cnt2 [26]int
	for i := 0; i < l1; i++ {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
		if s1[i] != s2[i] {
			wrong_pos++
		}
	}
	if wrong_pos > 2 {
		return false
	}
	for i := 0; i < 26; i++ {
		if cnt1[i] != cnt2[i] {
			return false
		}
	}
	return true
}
