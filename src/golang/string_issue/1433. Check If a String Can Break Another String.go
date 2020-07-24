package string_issue

//Input: s1 = "abc", s2 = "xya"
//Output: true
//Explanation: "ayx" is a permutation of s2="xya" which can break to string_issue "abc" which is a permutation of s1="abc".

func is_break(s1 string,s2 string)bool{
	l := len(s1)
	s1_large_num := 0
	s2_large_num := 0
	for i := 0;i < l;i++{
		if s1[i] > s2[i]{
			s1_large_num++
		}else if s1[i] < s2[i]{
			s2_large_num++
		}else{
			s1_large_num++
			s2_large_num++
		}
	}
	return (s1_large_num == l || s2_large_num == l)
}

func perm(s1 string,s2 string,pos int)bool{
	l := len(s1)
	if pos >= l{
		return false
	}
	for begin := pos;begin < l;begin++{
		ss1 := []byte(s1)
		ss1[begin],ss1[pos] = ss1[pos],ss1[begin]
		s1 = string(ss1)

		if is_break(s1,s2){
			return true
		}
		if perm(s1,s2,pos + 1){
			return true
		}
		ss1[begin],ss1[pos] = ss1[pos],ss1[begin]
		s1 = string(ss1)
	}
	return false
}

func CheckIfCanBreak(s1 string, s2 string) bool {
	return perm(s1,s2,0) || perm(s2,s1,0)
}

func CheckIfCanBreak2(s1 string, s2 string) bool {
	var m1 [26]int
	var m2 [26]int
	l := len(s1)
	for i := 0;i < l;i++{
		m1[s1[i] - 'a']++
		m2[s2[i] - 'a']++
	}

	s1_large_num := 0
	s2_large_num := 0
	pos1 := 0
	pos2 := 0
	for pos1 < 26 && pos2 < 26 {
		if m1[pos1] == 0{
			pos1++
			continue
		}
		if m2[pos2] == 0{
			pos2++
			continue
		}
		if pos1 < pos2{
			s1_large_num++
			m1[pos1]--
			m2[pos2]--
		}else if pos1 > pos2{
			s2_large_num++
			m2[pos2]--
			m1[pos1]--
		}else{
			s1_large_num++
			m1[pos1]--
			s2_large_num++
			m2[pos2]--
		}
	}
	return (s1_large_num == l || s2_large_num == l)
}