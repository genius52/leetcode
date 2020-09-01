package string_issue

import "strings"

//s = "abpcplea", d = ["ale","apple","monkey","plea"]
//
//Output:
//"apple"
func is_sub(s string,sub string)bool{
	var l int = len(s)
	var sub_len int = len(sub)
	var j int = 0
	for i := 0;i < l;i++{
		if s[i] == sub[j]{
			j++
		}
		if j == sub_len{
			return true
		}
	}
	return false
}

func FindLongestWord(s string, d []string) string {
	var res string
	var l int = len(s)
	var max_len int = 0
	for _,sub := range d{
		var sub_len int = len(sub)
		if l < sub_len || sub_len < max_len{
			continue
		}
		if is_sub(s,sub){
			if max_len < sub_len{
				res = sub
				max_len = sub_len
			}else if max_len == sub_len{
				if strings.Compare(sub,res) < 0{
					res = sub
				}
			}
		}
	}
	return res
}