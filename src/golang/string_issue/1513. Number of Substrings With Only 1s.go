package string_issue

//Input: s = "0110111"
//Output: 9
//Explanation: There are 9 substring in total with only 1's characters.
//"1" -> 5 times.
//"11" -> 3 times.
//"111" -> 1 time.
func NumSub(s string) int {
	s += "0"
	var l int = len(s)
	var res int = 0
	var begin int = 0
	var end int = 0
	for end < l{
		if s[end] == '1'{
			end++
		}else{
			cur_len := end - begin
			if cur_len > 0{
				res += (cur_len + 1) * cur_len /2
				res %= 1000000007
			}
			end++
			begin = end
		}
	}
	return res
}