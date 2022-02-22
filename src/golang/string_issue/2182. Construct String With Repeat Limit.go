package string_issue

import "strings"

//输入：s = "aababab", repeatLimit = 2
//输出："bbabaa"
func RepeatLimitedString(s string, repeatLimit int) string {
	var record [26]int
	for _,c := range s{
		record[c - 'a']++
	}
	var last int = -1
	var buf strings.Builder
	var types int = 0
	for i := 0;i <= 25;i++{
		if record[i] != 0{
			types++
		}
	}
	for true{
		var idx int = 25
		for idx > 0 && record[idx] == 0{
			if idx == last{
				idx--
				continue
			}
			idx--
			continue
		}
		if idx < 0 || idx == last{
			break
		}
		var add bool = false
		cnt := min_int(repeatLimit,record[idx])
		for i := 0;i < cnt;i++{
			add = true
			buf.WriteByte(byte('a' + idx))
			record[idx]--
		}
		if !add{
			break
		}
		if add && record[idx] == 0{
			last = idx
			types--
			continue
		}
		//add second biggest
		add = false
		for i := idx - 1;i >= 0;i--{
			if record[i] == 0{
				continue
			}
			add = true
			last = i
			if types == 1{
				cnt2 := min_int(repeatLimit,record[i])
				for j := 0;j < cnt2;j++{
					buf.WriteByte(byte('a' + i))
					record[i]--
				}
			}else{
				buf.WriteByte(byte('a' + i))
				record[i]--
			}
			if add && record[i] == 0{
				types--
			}
			break
		}
		if !add{
			break
		}
	}
	return buf.String()
}