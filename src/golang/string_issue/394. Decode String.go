package string_issue

import "strconv"

//394
//s = "3[a]2[bc]", return "aaabcbc".
//s = "3[a2[c]]", return "accaccacc".
//s = "2[abc]3[cd]ef", return "abcabccdcdcdef".
func recursive_decodeString(s string,l int,cur_pos *int)string{
	if *cur_pos >= l{
		return ""
	}
	var cur_str string
	//var num_begin int = *cur_pos
	var i int = *cur_pos
	for i < l{
		if s[i] == '['{
			var num_begin int = i - 1
			for num_begin >= 0{
				if s[num_begin] >= '0' && s[num_begin] <= '9'{
					num_begin--
				}else{
					break
				}
			}
			cnt,err := strconv.Atoi(s[num_begin + 1:i])
			if err == nil{
				var next_pos int = i + 1
				sub := recursive_decodeString(s,l,&next_pos)
				//begin = next_pos
				i = next_pos
				for cnt > 0{
					cur_str += sub
					cnt--
				}
			}
		}else if s[i] == ']'{
			*cur_pos = i + 1
			return cur_str
		}else if s[i] >= 'a' && s[i] <= 'z'{
			cur_str += string(s[i])
			i++
		}else{//digit
			i++
		}
	}
	return cur_str
}

func DecodeString2(s string) string{
	var i int = 0
	var l int = len(s)
	return recursive_decodeString(s,l,&i)
}

func decodeString(s string) string {
	var word_record []string
	var cnt_record []int
	var visit int = 0
	var word_num int
	var cur_string string
	for visit < len(s){
		if s[visit] == '['{
			cnt_record = append(cnt_record, word_num)
			word_record = append(word_record,cur_string)
			cur_string = ""
			word_num = 0
		}else if s[visit]==']'{
			pre := word_record[len(word_record) - 1]
			cnt := cnt_record[len(cnt_record) - 1]
			for i := 0;i < cnt;i++{
				pre += cur_string
			}
			cur_string = pre
			cnt_record = cnt_record[:len(word_record) - 1]
			word_record = word_record[:len(word_record) - 1]
		}else{
			if s[visit] < '0' || s[visit] > '9'{
				cur_string += string(s[visit])
			}else{
				c,_ := strconv.Atoi(string(s[visit]))
				word_num = word_num * 10 + c
			}
		}
		visit++
	}
	return cur_string
}