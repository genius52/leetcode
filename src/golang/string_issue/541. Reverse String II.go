package string_issue

//Input: s = "abcdefg", k = 2
//Output: "bacdfeg"
func ReverseStr(s string, k int) string {
	if k == 1{
		return s
	}
	var l int = len(s)
	var res string
	var index int = 0
	var reverse bool = true
	for index < l{
		if reverse{
			for i := min_int(l - 1,index + k - 1);i >= index;i--{
				res += string(s[i])
			}
		}else{
			for i := index;i < l && i < index + k;i++{
				res += string(s[i])
			}
		}
		reverse = !reverse
		index += k
	}
	return res
}