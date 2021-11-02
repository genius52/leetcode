package string_issue

//Input: s = "aacaba"
//Output: 2
//Explanation: There are 5 ways to split "aacaba" and 2 of them are good.
//("a", "acaba") Left string and right string contains 1 and 3 different letters respectively.
//("aa", "caba") Left string and right string contains 1 and 3 different letters respectively.
//("aac", "aba") Left string and right string contains 2 and 2 different letters respectively (good split).
//("aaca", "ba") Left string and right string contains 2 and 2 different letters respectively (good split).
//("aacab", "a") Left string and right string contains 3 and 1 different letters respectively.
func numSplits(s string) int{
	var l int = len(s)
	var res int = 0
	var record [26]int
	for i := 0;i < l;i++{
		record[s[i] - 'a']++
	}
	var other_kinds int = 0
	var cur [26]int
	for i := 0;i < 26;i++{
		if record[i] != 0{
			other_kinds++
		}
	}
	var cur_kinds int = 0
	for i := 0;i < l;i++{
		c := s[i] - 'a'
		if cur[c] == 0{
			cur_kinds++

		}
		cur[c]++
		if record[c] == cur[c]{
			other_kinds--
		}
		if cur_kinds == other_kinds{
			res++
		}
		if cur_kinds > other_kinds{
			break
		}
	}
	return res
}

func NumSplits(s string) int {
	var l int = len(s)
	var dp_left []int = make([]int,l)
	var dp_right []int = make([]int,l)
	var record_left map[uint8]int = make(map[uint8]int)
	var record_right map[uint8]int = make(map[uint8]int)
	for i := 0;i < l;i++{
		record_left[s[i]]++
		dp_left[i] = len(record_left)
		record_right[s[l - 1 - i]]++
		dp_right[l - 1 - i] = len(record_right)
	}
	var res int = 0
	for i := 0;i < l - 1;i++{
		if(dp_left[i] == dp_right[i + 1]){
			res++
		}
	}
	return res
}
