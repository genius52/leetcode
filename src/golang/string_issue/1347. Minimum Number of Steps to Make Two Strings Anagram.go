package string_issue

func minSteps2(s string, t string) int{
	var record_s [26]int
	for _,c := range s{
		record_s[c - 'a']++
	}
	for _,c := range t{
		record_s[c - 'a']--
	}
	var res int = 0
	for i := 0;i < 26;i++{
		if record_s[i] > 0{
			res += record_s[i]
		}
	}
	return res
}

//Input: s = "leetcode", t = "practice"
//Output: 5
//Explanation: Replace 'p', 'r', 'a', 'i' and 'c' from t with proper characters to make t anagram of s.
func minSteps(s string, t string) int {
	var record map[int32]int = make(map[int32]int)
	for _,c := range s{
		if _,ok := record[c];ok{
			record[c]++
		}else{
			record[c] = 1
		}
	}
	for _,c := range t{
		if _,ok := record[c];ok{
			record[c]--
			if record[c] == 0{
				delete(record,c)
			}
		}
	}
	var res int = 0
	for _,l := range record{
		res += l
	}
	return res
}