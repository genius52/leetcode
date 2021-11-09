package string_issue

//Input: s = "??yw?ipkj?"
//Output: "acywaipkja"
func modifyString(s string) string {
	var l int = len(s)
	backup := [3]uint8{'a','b','c'}
	var res string
	for i := 0;i < l;i++{
		if s[i] != '?'{
			res += string(s[i])
		}else{
			for j := 0;j < 3;j++{
				if (i == 0 && i == l - 1) || (i == 0 && backup[j] != s[i + 1]) || (i == l - 1 && backup[j] != res[i - 1]) || (i > 0 && i < l - 1 && backup[j] != res[i - 1] && backup[j] != s[i + 1]){
					res += string(backup[j])
					break
				}
			}
		}
	}
	return res
}

func ModifyString(s string) string {
	var res string
	s = "0" + s
	var l int = len(s)
	for i := 0;i < l;i++{
		if s[i] != '?'{
			res += string(s[i])
		}else{
			var last uint8 = res[i - 1]
			var next uint8 = 0
			if i != l - 1{
				next = s[i + 1]
			}
			var c uint8 = 'a'
			for c == last || c == next{
				c++;
				if c > 'z'{
					c = 'a';
				}
			}
			res += string(c)
		}
	}
	res = res[1:]
	return res
}