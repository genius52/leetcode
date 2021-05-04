package string_issue

import "strconv"

func dp_splitString(s string,l int,cur_pos int,prev_num int)bool{
	if cur_pos >= l{
		return false
	}
	num1,err1 := strconv.Atoi(s[cur_pos:])
	if err1 == nil{
		if (prev_num - num1) == 1{
			return true
		}
	}
	var res bool = false
	for i := cur_pos + 1;i < l;i++{
		num,err := strconv.Atoi(s[cur_pos:i])
		if err == nil{
			if (prev_num - num) == 1{
				res = dp_splitString(s,l,i,num)
				if res{
					return true
				}
			}else if prev_num < num{
				break
			}
		}
	}
	return res
}

func SplitString(s string) bool {
	var l int = len(s)
	for i := 1;i < l;i++{
		num,err := strconv.Atoi(s[:i])
		if err == nil{
			res := dp_splitString(s,l,i,num)
			if res{
				return true
			}
		}
	}
	return false
}