package string_issue

import "sort"

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

func CheckIfCanBreak3(s1 string, s2 string) bool{
	var l int = len(s1)
	if l == 1{
		return true
	}
	b1 := []byte(s1)
	sort.Slice(b1, func(i, j int) bool {
		return b1[i] <= b1[j]
	})
	b2 := []byte(s2)
	sort.Slice(b2, func(i, j int) bool {
		return b2[i] <= b2[j]
	})
	if l == 2{
		return (b1[0] <= b2[0] && b1[1] <= b2[1]) || (b1[0] >= b2[0] && b1[1] >= b2[1])
	}
	var meet1 bool = true
	for i := 1;i < l - 1;i++{
		if (b1[i - 1] <= b2[i - 1] && b1[i] <= b2[i] && b1[i + 1] <= b2[i + 1]){
			continue
		}else{
			meet1 = false
			break
		}
	}
	var meet2 bool = true
	for i := 1;i < l - 1;i++{
		if (b1[i - 1] >= b2[i - 1] && b1[i] >= b2[i] && b1[i + 1] >= b2[i + 1]){
			continue
		}else{
			meet1 = false
			break
		}
	}
	return meet1 || meet2
}