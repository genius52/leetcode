package string_issue

//Input: s = "aba"
//Output: 6
//Explanation: The 6 distinct subsequences are "a", "b", "ab", "ba", "aa" and "aba".

func DistinctSubseqII(s string) int {
	var l int = len(s)
	var record []int = make([]int,l)
	var res int = 0
	for i := 0;i < l;i++{
		record[i] = 1
		for j := i - 1;j >= 0;j--{
			if s[i] != s[j]{
				record[i] += record[j]
				record[i] %= (1e9 + 7)
			}
		}
		res += record[i]
		res %= (1e9 + 7)
	}
	return res
}
//func DistinctSubseqII(s string) int{
//	var record [26]int
//	var l int = len(s)
//	for i := 0;i < l;i++{
//		var cur int = 0
//		for j := 0;j < 26;j++{
//			if record[j] == 0{
//				continue
//			}
//			cur += record[j]
//		}
//		if record[int(s[i] - 'a')] == 0{
//			record[s[i] - 'a'] = cur + 1
//		}else{
//			record[s[i] - 'a'] += cur
//		}
//	}
//	var res int = 0
//	for i := 0;i < 26;i++{
//		res += record[i]
//		res = res % (1e9 + 7)
//	}
//	return res
//}

//TLE
//func DistinctSubseqII(s string) int {
//	var l int = len(s)
//	var record [26]map[string]bool
//	for i := 0;i < 26;i++{
//		record[i] = make(map[string]bool)
//	}
//	for i := 0;i < l;i++{
//		var cur map[string]bool = make(map[string]bool)
//		for j := 0;j < 26;j++{
//			for sub,_ := range record[j]{
//				s2 := sub + string(s[i])
//				cur[s2] = true
//			}
//		}
//		cur[string(s[i])] = true
//		record[s[i] - 'a'] = cur
//	}
//	var res int = 0
//	for i := 0;i < 26;i++{
//		res += len(record[i])
//		res = res % (1e9 + 7)
//	}
//	return res
//}