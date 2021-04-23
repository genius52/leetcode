package string_issue

//Input: s1 = "ab" s2 = "eidbaooo"
//Output: True
//Explanation: s2 contains one permutation of s1 ("ba").
func is_equal(dict1 [26]int,dict2 [26]int)bool{
	for i := 0;i < 26;i++{
		if dict1[i] != dict2[i]{
			return false
		}
	}
	return true
}

func CheckInclusion(s1 string, s2 string) bool {
	var l1 int = len(s1)
	var l2 int = len(s2)
	if l1 > l2 {
		return false
	}
	var dict1 [26]int
	var dict2 [26]int
	for i := 0;i < l1;i++{
		dict1[s1[i] - 'a']++
		dict2[s2[i] - 'a']++
	}
	if is_equal(dict1,dict2){
		return true
	}
	for i := l1;i < l2;i++{
		dict2[s2[i - l1] - 'a']--
		dict2[s2[i] - 'a']++
		if is_equal(dict1,dict2){
			return true
		}
	}
	return false
}