package string_issue

//Input: s = "aaaabbbbcccc"
//Output: "abccbaabccba"
func SortString(s string) string{
	var record [26]int
	var begin int = 26
	var end int = 0
	for _,c := range s{
		pos := int(c - 'a')
		record[pos]++
		if pos < begin{
			begin = pos
		}
		if pos > end{
			end = pos
		}
	}
	if begin == end{
		return s
	}
	var res string
	var idx int = 0
	var l int = len(s)
	for idx < l{
		for i := begin;i <= end;i++{
			if record[i] > 0{
				res += string(i + 'a')
				idx++
				record[i]--
			}
		}
		for i := end;i >= begin;i--{
			if record[i] > 0{
				res += string(i + 'a')
				idx++
				record[i]--
			}
		}
	}
	return res
}

//func sortString(s string) string {
//	var record [26]int
//	for _,c := range s{
//		record[c-'a']++
//	}
//	var head_to_tail bool = true
//	var res string
//	head,tail := 0,25;
//	for head <= tail {
//		if record[head] == 0{
//			head++
//			continue
//		}
//		if record[tail] == 0{
//			tail--
//			continue
//		}
//		i := head
//		j := tail
//		for i <= j{
//			if head_to_tail{
//				if record[i] == 0{
//					i++
//					continue
//				}
//			}else{
//				if record[j] == 0{
//					j--
//					continue
//				}
//			}
//			if head_to_tail{
//				res += string(i + 'a')
//				record[i]--
//				i++
//			}else{
//				res += string(j + 'a')
//				record[j]--
//				j--
//			}
//		}
//		head_to_tail = !head_to_tail
//	}
//	return res
//}