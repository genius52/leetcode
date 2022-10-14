package string_issue

//Input: s = "abab"
//Output: "bab"
//Explanation: The substrings are ["a", "ab", "aba", "abab", "b", "ba", "bab"].
//The lexicographically maximum substring is "bab".
func LastSubstring(s string) string{
	var l int = len(s)
	var left int = 0
	var right int = 1
	var cur_len int = 0
	for right + cur_len < l{
		if s[left + cur_len] == s[right + cur_len]{
			cur_len++
		}else if s[left + cur_len] < s[right + cur_len]{
			//left = right
			//cur_len = 0
			//right++
			//i=max(i+k+1,j),j=i+1,k=0;
			left = max_int(right,left + cur_len + 1)
			right = left + 1
			cur_len = 0
		}else if s[left + cur_len] > s[right + cur_len]{
			right += cur_len + 1
			cur_len = 0
		}
	}
	return s[left:]
}

//func lastSubstring(s string) string {
//	var l int = len(s)
//	i := 0
//	j := 1
//	k := 0
//	for (j+k) < l{
//		if s[i+k] == s[j+k]{//i和j相同，继续比较下一个
//			k++
//		}else if s[i+k] < s[j+k]{
//			i = max_int(i + k + 1,j)
//			j = i + 1
//			k = 0 //长度重置为0
//		}else{
//			j += k + 1//？？？
//			k = 0
//		}
//	}
//	return s[i:]
//}
//
//func lastSubstring2(s string) string {
//	var i int = 0
//	var j int = 1
//	var k int = 0
//	var l int = len(s)
//	for (j + k) < l{
//		if s[i + k] == s[j + k]{
//			k++
//		}else if s[i + k] > s[j + k]{
//			if (i + k + 1) > j{
//				i = i + k + 1
//
//			}else{
//				i = j
//			}
//			j = i + 1
//			k = 0
//		}else if s[i + k] < s[j + k]{
//			i = j
//			j++
//			k = 0
//		}
//	}
//	return s[i:]
//}