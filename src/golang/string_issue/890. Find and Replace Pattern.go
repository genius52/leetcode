package string_issue

import "strconv"

//输入：words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
//输出：["mee","aqq"]

func FindAndReplacePattern(words []string, pattern string) []string {
	var res []string
	var l int = len(pattern)
	for _,word := range words{
		var origin_cur map[uint8]uint8 = make(map[uint8]uint8)
		var cur_origin map[uint8]uint8 = make(map[uint8]uint8)
		var same bool = true
		for i := 0;i < l;i++{
			if cur_c,ok1 := origin_cur[pattern[i]];!ok1{
				if _,ok2 := cur_origin[word[i]];ok2{
					same = false
					break
				}
				origin_cur[pattern[i]] = word[i]
				cur_origin[word[i]] = pattern[i]
			}else{
				if origin_c,ok2 := cur_origin[word[i]];!ok2{
					same = false
					break
				}else{
					if !(cur_c == word[i]) || !(origin_c == pattern[i]){
						same = false
						break
					}
				}
			}
		}
		if same{
			res = append(res,word)
		}
	}
	return res
}

func findAndReplacePattern(words []string, pattern string) []string {
	var res []string
	target := convert_findAndReplacePattern(pattern)
	for _,w := range words{
		if convert_findAndReplacePattern(w) == target{
			res = append(res,w)
		}
	}
	return res
}

func convert_findAndReplacePattern(s string)string{
	var res string
	var record map[int32]int = make(map[int32]int)
	var cnt int = 0
	for _,c := range s{
		if val,ok := record[c];ok{
			res += strconv.Itoa(val)
		}else{
			record[c] = cnt
			res += strconv.Itoa(cnt)
			cnt++
		}
	}
	return res
}